package database

import (
	"ChipiAiAuth/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func InitDb() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	dns := os.Getenv("DB_DNS_AUTH")
	if dns == "" {
		log.Println("DB_DNS_AUTH environment variable not set")
	}

	var err error
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	log.Println("Database connection established")

	if err != nil {
		log.Println(err)
	}

	if err := Db.AutoMigrate(&models.User{}); err != nil {
		log.Println(err)
	}
}
