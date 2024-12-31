package email_engine

import (
	"context"
	"encoding/csv"
	"errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/fs"
	"log"
	"os"
	"reflect"
	"strconv"
	"summarize-transactions/config"
	"summarize-transactions/core"
	"summarize-transactions/models"
	"summarize-transactions/services"
)

type EmailEngine struct {
	*services.IMAPClient
	*config.StorageConfig
	LocalFolder string
}

func Initialize() (*EmailEngine, error) {
	engineConfig := config.NewEmailEngineConfig()
	imapConfig := engineConfig.ImapConfig
	storageConfig := engineConfig.StorageConfig
	localFolder := engineConfig.LocalFolder

	c, err := services.NewIMAPClient(imapConfig)

	if err != nil {
		return nil, err
	}

	return &EmailEngine{c, &storageConfig, localFolder}, nil
}

func (eng *EmailEngine) Run() {
	defer func() {
		if logoutErr := eng.Logout(); logoutErr != nil {
			log.Fatalf("Error during logout: %v", logoutErr)
		}
		log.Printf("Sucessfully logget out from imap client")
	}()

	if _, err := eng.Select("Inbox", false); err != nil {
		log.Fatal(err)
	}

	done := make(chan error, 1)
	msgch := make(chan *services.EmailMessage, 10)

	go func() {
		subject := "Extrato da fatura do Cartão Nubank"
		done <- eng.FetchMessages(subject, msgch)
	}()

	for msg := range msgch {
		engMessage := NewEmailEngineMessage(msg)

		if err := eng.ProcessMessage(engMessage); err != nil {
			log.Fatal(err)
		}
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}

func (eng *EmailEngine) saveFileLocally(localFilePath string, bodyData []byte) error {
	return os.WriteFile(localFilePath, bodyData, fs.ModePerm)
}

func (eng *EmailEngine) uploadFileToStorage(file *os.File, fileName string) error {
	ctx := context.Background()
	objectName := eng.EmailCsvFolder + "/" + fileName
	storageClient := services.NewStorageClient(ctx)
	return storageClient.Upload(file, eng.BucketName, objectName)
}

func (eng *EmailEngine) ProcessMessage(engMsg *EmailEngineMessage) error {
	fileName := engMsg.FileName
	b := engMsg.AttachmentBody
	localFilePath := eng.LocalFolder + "/" + fileName

	if err := eng.saveFileLocally(localFilePath, b); err != nil {
		return err
	}

	localFile, err := os.Open(localFilePath)

	if err != nil {
		return err
	}

	if err = eng.uploadFileToStorage(localFile, fileName); err == nil {
		return err
	}

	return nil
}

func outrasCoisas() {
	//filename := "extract.csv"
	//loadCsvToDb(filename)

	categorizedFilename := "extract.csv"
	saveTransactionsCategoriesFromCsv(categorizedFilename)
}

func loadCsvToDb(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Println("Error opening file: ", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing file: ", err)
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Println("Error reading records from file: ", err)
	}

	csvMap := mapCsvRecordsToMap(records)

	var transactions []models.Transaction
	for _, csvItem := range csvMap {
		transaction := models.Transaction{}
		for key, value := range csvItem {
			caser := cases.Title(language.English)
			titledKey := caser.String(key)
			structField, _ := reflect.TypeOf(transaction).FieldByName(titledKey)
			category := structField.Tag.Get("csv")
			if category != "" {
				reflect.ValueOf(&transaction).Elem().FieldByName(titledKey).Set(reflect.ValueOf(value))
			}
		}

		amount := csvItem["amount"]
		parsedAmount, err := strconv.ParseFloat(amount.(string), 64)

		if err != nil {
		}

		transaction.Amount = parseMoneyFloatToCurrency(parsedAmount)

		transactions = append(transactions, transaction)
	}

	log.Println("Finished parsing CSV to Transactions, inserting to database...")

	insertTransactionsToDb(transactions)
}

func connectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func insertTransactionsToDb(transactions []models.Transaction) {
	db, err := connectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	result := db.Create(&transactions)

	if result.Error != nil {
		log.Printf("failed to insert data: %v", result.Error)
	}

	log.Printf("successfully inserted transactions to database")
}

func saveTransactionsCategoriesFromCsv(filename string) {
	db, err := connectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing file: ", err)
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Println("Error reading records from file: ", err)
	}

	csvMap := mapCsvRecordsToMap(records)

	var categories []models.Category
	result := db.Find(&categories)

	if result.Error != nil {
		log.Fatal("Failed to retrieve categories:", result.Error)
	}

	for _, csvItem := range csvMap {
		var transaction models.Transaction
		csvTitle := csvItem["title"]
		csvAmount, _ := strconv.ParseFloat(csvItem["amount"].(string), 64)
		csvAmountInt := parseMoneyFloatToCurrency(csvAmount)
		csvDate := csvItem["date"]
		csvCategory := csvItem["category"].(string)

		result = db.Where("title = ? AND amount = ? AND date = ?", csvTitle, csvAmountInt, csvDate).First(&transaction)

		category, categoryErr := findCategoryByName(categories, csvCategory)

		if categoryErr != nil {
			log.Printf("Failed to retrieve category:", categoryErr)
			continue
		}

		// TODO: cuidado pra não fazer append duplicado!
		transaction.Categories = append(transaction.Categories, category)

		if result.Error != nil {
			log.Printf("Failed to retrieve transaction: %e", result.Error)
			continue
		}

		result = db.Save(&transaction)

		if result.Error != nil {
			log.Printf("Failed to save transaction: %e", result.Error)
		}
	}
}

func parseMoneyFloatToCurrency(floatNum float64) core.Currency {
	return core.Currency(floatNum * 100)
}

func mapCsvRecordsToMap(records [][]string) []map[string]interface{} {
	var csvMap []map[string]interface{}

	// Use the first row as header (column names)
	headers := records[0]

	// Iterate over the records (starting from the second row)
	for _, record := range records[1:] {
		rowMap := make(map[string]interface{})

		for i, value := range record {
			// Use reflection to dynamically set map keys
			rowMap[headers[i]] = value
		}
		csvMap = append(csvMap, rowMap)
	}

	return csvMap
}

func findCategoryByName(categories []models.Category, term string) (models.Category, error) {
	var categoryToReturn models.Category
	var err error = nil

	for _, category := range categories {
		if category.Name == term {
			categoryToReturn = category
			break
		}
	}

	if categoryToReturn.Name == "" {
		err = errors.New("category not found")
	}

	return categoryToReturn, err
}
