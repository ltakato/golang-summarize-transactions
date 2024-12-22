package main

import (
	"encoding/csv"
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

				switch h := p.Header.(type) {
				case *mail.AttachmentHeader:
					filename, _ := h.Filename()

					log.Printf("Got attachment: %v", filename)

					b, errp := io.ReadAll(p.Body)

					if errp != nil {
						fmt.Println("errp ===== :", errp)
					}

					err := os.WriteFile(filename, b, fs.ModePerm)

					if err != nil {
						log.Println("attachment err: ", err)
					}

					loadCsvToDb(filename)
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

		transaction.Amount = int(parsedAmount * 100)

		transactions = append(transactions, transaction)
	}

	log.Println("Finished parsing CSV to Transactions, inserting to database...")

	insertTransactionsToDb(transactions)
}

func insertTransactionsToDb(transactions []models.Transaction) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	result := db.Create(&transactions)

	if result.Error != nil {
		log.Printf("failed to insert data: %v", result.Error)
	}

	log.Printf("successfully inserted transactions to database")
}
