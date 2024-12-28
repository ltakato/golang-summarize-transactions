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

	dummyCategories := generateDummyCategories()
	result := db.Create(&dummyCategories)

	if result.Error != nil {
		log.Fatalf("failed to insert categories: %v", err)
	}

	//insertCategoryTerms(db, dummyCategory1, dummyCategory2)

	insertTransactions(db, dummyCategories)

	fmt.Println("Dummy data loaded successfully!")
}

func generateDummyCategories() []*models.Category {
	dummyCategory1 := models.Category{
		Name: "refeição",
	}
	dummyCategory2 := models.Category{
		Name: "transporte",
	}
	dummyCategory3 := models.Category{
		Name: "assinatura",
	}

	return []*models.Category{&dummyCategory1, &dummyCategory2, &dummyCategory3}
}

func insertCategoryTerms(db *gorm.DB, dummyCategory1, dummyCategory2 models.Category) {
	dummyCategoryTerm1 := models.CategoryTerms{
		Expression: "restaurante",
		Category:   dummyCategory1,
	}
	dummyCategoryTerm2 := models.CategoryTerms{
		Expression: "outback",
		Category:   dummyCategory2,
	}
	dummyCategoryTerm3 := models.CategoryTerms{
		Expression: "uber",
		Category:   dummyCategory2,
	}

	db.Create([]*models.CategoryTerms{&dummyCategoryTerm1, &dummyCategoryTerm2, &dummyCategoryTerm3})
}

func insertTransactions(db *gorm.DB, dummyCategories []*models.Category) {
	dummyCategory1 := *dummyCategories[0]
	dummyCategory2 := *dummyCategories[1]

	dummyTransaction1 := models.Transaction{
		Title:      "Restaurante Reino das Carnes",
		Date:       "2024-12-10",
		Amount:     8000,
		Categories: []models.Category{dummyCategory1},
	}
	dummyTransaction2 := models.Transaction{
		Title:  "Restaurante Frango Frito e Assado",
		Date:   "2024-12-15",
		Amount: 12835,
	}
	dummyTransaction3 := models.Transaction{
		Title:      "Outback Steakhouse",
		Date:       "2024-12-28",
		Amount:     58095,
		Categories: []models.Category{dummyCategory1},
	}
	dummyTransaction4 := models.Transaction{
		Title:      "Outback Steakhouse",
		Date:       "2024-11-03",
		Amount:     8000,
		Categories: []models.Category{dummyCategory1},
	}
	dummyTransaction5 := models.Transaction{
		Title:      "Uber* Trip",
		Date:       "2024-12-28",
		Amount:     3000,
		Categories: []models.Category{dummyCategory2},
	}
	dummyTransaction6 := models.Transaction{
		Title:      "Uber* Trip",
		Date:       "2024-12-28",
		Amount:     4500,
		Categories: []models.Category{dummyCategory2},
	}
	dummyTransaction7 := models.Transaction{
		Title:      "Uber* Trip",
		Date:       "2024-11-03",
		Amount:     5397,
		Categories: []models.Category{dummyCategory2},
	}

	transactions := []*models.Transaction{&dummyTransaction1, &dummyTransaction2, &dummyTransaction3, &dummyTransaction4, &dummyTransaction5, &dummyTransaction6, &dummyTransaction7}

	result := db.Create(&transactions)

	if result.Error != nil {
		log.Printf("failed to insert transactions: %v", result.Error)
	}
}
