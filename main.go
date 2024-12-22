package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"io/fs"
	"log"
	"net/textproto"
	"os"
	"reflect"
	"strconv"
	"summarize-transactions/models"
)

func main() {
	filename := "extract.csv"
	saveCsvToFile(filename)
	loadCsvToDb(filename)

	categorizedFilename := "extract.csv"
	saveTransactionsCategoriesFromCsv(categorizedFilename)
}

func parseMoneyFloatToInt(floatNum float64) int {
	return int(floatNum * 100)
}

func saveCsvToFile(filename string) {
	// connect client
	imapServer := os.Getenv("IMAP_SERVER")
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASSWORD")
	c, err := client.DialTLS(imapServer, nil)

	if err != nil {
		log.Fatal(err)
	}

	// defer logout
	defer func() {
		if logoutErr := c.Logout(); logoutErr != nil {
			log.Printf("Error during logout: %v", logoutErr)
		}
	}()

	// Login
	if err := c.Login(email, password); err != nil {
		log.Fatal(err)
	}

	// Select mailbox
	_, err = c.Select("Inbox", false)

	if err != nil {
		log.Fatal(err)
	}

	// Define search criteria
	nubankExtractSubject := "Extrato da fatura do Cartão Nubank"
	criteria := imap.SearchCriteria{
		Header: textproto.MIMEHeader{"Subject": {nubankExtractSubject}},
	}

	// Perform the search
	seqNums, err := c.Search(&criteria)

	done := make(chan error, 1)

	if len(seqNums) > 0 {
		// Fetch matching messages
		seqset := new(imap.SeqSet)
		seqset.AddNum(seqNums...)

		section := &imap.BodySectionName{}
		messages := make(chan *imap.Message, 10)

		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
		}()

		for msg := range messages {
			log.Println("* " + msg.Envelope.Subject)

			mr, err := mail.CreateReader(msg.GetBody(section))

			if err != nil {
				log.Fatal(err)
			}

			for {
				p, err := mr.NextPart()

				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch p.Header.(type) {
				case *mail.AttachmentHeader:
					log.Printf("Got attachment")

					b, errp := io.ReadAll(p.Body)

					if errp != nil {
						fmt.Println("failed to read attachment body", errp)
					}

					err := os.WriteFile(filename, b, fs.ModePerm)

					if err != nil {
						log.Println("saving attachment file err: ", err)
					}
				}
			}
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		log.Println("Done!")
	}
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
			tag := structField.Tag.Get("csv")
			if tag != "" {
				reflect.ValueOf(&transaction).Elem().FieldByName(titledKey).Set(reflect.ValueOf(value))
			}
		}

		amount := csvItem["amount"]
		parsedAmount, err := strconv.ParseFloat(amount.(string), 64)

		if err != nil {
		}

		transaction.Amount = parseMoneyFloatToInt(parsedAmount)

		transactions = append(transactions, transaction)
	}

	log.Println("Finished parsing CSV to Transactions, inserting to database...")

	insertTransactionsToDb(transactions)
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

	var tags []models.Tag
	result := db.Find(&tags)

	if result.Error != nil {
		log.Fatal("Failed to retrieve tags:", result.Error)
	}

	for _, csvItem := range csvMap {
		var transaction models.Transaction
		csvTitle := csvItem["title"]
		csvAmount, _ := strconv.ParseFloat(csvItem["amount"].(string), 64)
		csvAmountInt := parseMoneyFloatToInt(csvAmount)
		csvDate := csvItem["date"]
		csvTag := csvItem["tag"].(string)

		result = db.Where("title = ? AND amount = ? AND date = ?", csvTitle, csvAmountInt, csvDate).First(&transaction)

		tag, tagErr := findTagByName(tags, csvTag)

		if tagErr != nil {
			log.Printf("Failed to retrieve tag:", tagErr)
			continue
		}

		// TODO: cuidado pra não fazer append duplicado!
		transaction.Tags = append(transaction.Tags, tag)

		if result.Error != nil {
			log.Printf("Failed to retrieve transaction:", result.Error)
			continue
		}

		result = db.Save(&transaction)

		if result.Error != nil {
			log.Printf("Failed to save transaction:", result.Error)
		}
	}
}

func findTagByName(tags []models.Tag, term string) (models.Tag, error) {
	var tagToReturn models.Tag
	var err error = nil

	for _, tag := range tags {
		if tag.Name == term {
			tagToReturn = tag
			break
		}
	}

	if tagToReturn.Name == "" {
		err = errors.New("tag not found")
	}

	return tagToReturn, err
}
