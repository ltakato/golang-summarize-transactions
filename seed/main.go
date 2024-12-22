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

	dummyTag1 := models.Tag{
		Name: "refeições",
	}

	dummyTag2 := models.Tag{
		Name: "transporte",
	}

	result := db.Create([]*models.Tag{&dummyTag1, &dummyTag2})

	dummyTagTerm1 := models.TagTerms{
		Expression: "restaurante",
		Tag:        dummyTag1,
	}
	dummyTagTerm2 := models.TagTerms{
		Expression: "outback",
		Tag:        dummyTag1,
	}
	dummyTagTerm3 := models.TagTerms{
		Expression: "uber",
		Tag:        dummyTag2,
	}

	db.Create([]*models.TagTerms{&dummyTagTerm1, &dummyTagTerm2, &dummyTagTerm3})

	dummyTransaction1 := models.Transaction{
		Title:  "Restaurante Reino das Carnes",
		Date:   "2024-12-10",
		Amount: 8000,
		Tags:   []models.Tag{dummyTag1},
	}
	dummyTransaction2 := models.Transaction{
		Title:  "Restaurante Frango Frito e Assado",
		Date:   "2024-12-15",
		Amount: 12835,
		Tags:   []models.Tag{dummyTag1},
	}
	dummyTransaction3 := models.Transaction{
		Title:  "Outback Steakhouse",
		Date:   "2024-12-28",
		Amount: 58095,
		Tags:   []models.Tag{dummyTag1},
	}
	dummyTransaction4 := models.Transaction{
		Title:  "Outback Steakhouse",
		Date:   "2024-11-03",
		Amount: 8000,
		Tags:   []models.Tag{dummyTag1},
	}
	dummyTransaction5 := models.Transaction{
		Title:  "Uber* Trip",
		Date:   "2024-12-28",
		Amount: 3000,
		Tags:   []models.Tag{dummyTag2},
	}
	dummyTransaction6 := models.Transaction{
		Title:  "Uber* Trip",
		Date:   "2024-12-28",
		Amount: 4500,
		Tags:   []models.Tag{dummyTag2},
	}
	dummyTransaction8 := models.Transaction{
		Title:  "Uber* Trip",
		Date:   "2024-11-03",
		Amount: 5397,
		Tags:   []models.Tag{dummyTag2},
	}

	transactions := []*models.Transaction{&dummyTransaction1, &dummyTransaction2, &dummyTransaction3, &dummyTransaction4, &dummyTransaction5, &dummyTransaction6, &dummyTransaction8}

	result = db.Create(&transactions)

	if result.Error != nil {
		log.Printf("failed to insert data: %v", result.Error)
	}

	fmt.Println("Dummy data loaded successfully!")
}
