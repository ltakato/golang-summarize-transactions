package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"summarize-transactions/controllers"
	"summarize-transactions/core"
	"summarize-transactions/dto"
	"summarize-transactions/email_engine"
	"summarize-transactions/repositories"
)

func main() {
	initializeApi()
	//initializeEmailEngine()
}

func initializeEmailEngine() {
	engine, err := email_engine.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	engine.Run()
}

func initializeApi() {
	db, err := connectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	transactionsRepository := repositories.New(db)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("x-user-id")
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("partial_iso8601", func(fl validator.FieldLevel) bool {
			return core.IsValidPartialISO8601(fl.Field().String())
		})
		if err != nil {
			log.Fatalf("failed to register validation: %v", err)
		}
	}

	router.Use(cors.New(config))

	router.Use(UserIdMiddleware())

	categoriesController := controllers.NewCategoriesController(transactionsRepository)
	summaryController := controllers.NewSummaryController(transactionsRepository)

	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/summary", summaryController.GetSummary())
		categoryRouter := apiRouter.Group("/categories")
		categoryRouter.Use(CategoryQueryMiddleware())
		{
			categoryRouter.GET("", categoriesController.GetCategories())
			categoryRouter.GET("/:id/transactions", categoriesController.GetCategoryTransactions())
		}
	}

	err = router.Run(":8080")

	if err != nil {
		log.Fatalf("failed to run API: %v", err)
	}
}

func UserIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("x-user-id")

		if userId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "x-user-id header is required"})
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}

func CategoryQueryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var q dto.CategoryQuery
		err := c.ShouldBindWith(&q, binding.Query)

		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			c.Abort()
			return
		}

		c.Set("categoryQuery", q)

		c.Next()
	}
}

func connectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
