package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todobackend/internal/todoservice"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := DB.AutoMigrate(&todoservice.Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
	return DB
}
