package db

import (
	"fmt"
	"log"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/andrewdwi2198/air-aid/helpers/models"
	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	username := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")
	port := os.Getenv("PG_PORT")
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, pass, dbname)

	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Printf("[DB][ConnectDB] %s\n", err)
		return nil, err
	}
	log.Printf("[DB][ConnectDB] Successfully established connection to %s\n", dbname)

	db.AutoMigrate(&models.Data{})

	return db, nil
}
