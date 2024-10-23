package main

import (
	"customer-feedback-api/controllers"
	"customer-feedback-api/models"
	"customer-feedback-api/repositories"
	"customer-feedback-api/routes"
	"customer-feedback-api/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Get database connection settings from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize the database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		dbSQL.Close()
	}()

	// Disable logger
	db.Logger = logger.New(log.New(os.Stdout, "", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second * 2,
		LogLevel:      logger.Info,
		Colorful:      false,
	})

	// Auto-migrate the feedback table
	err = db.AutoMigrate(&models.Feedback{})
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	repository := repositories.NewFeedbackRepository(db)

	// Initialize services
	service := services.NewFeedbackService(repository)

	// Initialize controllers
	controller := controllers.NewFeedbackController(service)

	// Initialize routes
	routes.InitRoutes(http.DefaultServeMux, controller)
	handler := corsMiddleware(http.DefaultServeMux)
	// Start the HTTP server
	log.Println("Starting the HTTP server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
