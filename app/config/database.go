package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	godotenv.Load(".env")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbName, password)
	db, err := gorm.Open(postgres.Open(conname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}

	return db
}
