package main

import (
	"log"

	"github.com/203ZK/gofit-backend/internal/config"
	"github.com/203ZK/gofit-backend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	cfg, err := config.LoadConfig("dbconfig.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	err = database.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Database connection successful")
}
