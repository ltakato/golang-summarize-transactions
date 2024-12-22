package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"summarize-transactions/models"
)

func main() {
	fmt.Println("seed called")

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	dummyTag := models.Tag{
		Name: "Meals",
	}

	dummyTransaction := models.Transaction{
		Title: "Test Transaction",
		Date:  "2024-12-10",
		Tags:  []models.Tag{dummyTag},
	}

	err = db.Create(&dummyTransaction).Error

	if err != nil {
		log.Printf("failed to insert data: %v", err)
	}

	fmt.Println("Dummy data loaded successfully!")
}
