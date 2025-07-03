package main

import (
	"log"
	"os"

	"github.com/203ZK/gofit-backend/internal/config"
	"github.com/203ZK/gofit-backend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing command")
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	cfg, err := config.LoadConfig("dbconfig.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	switch os.Args[1] {
	case "migrateDBs":
		database.MigrateDBs(db)
	case "dropDBs":
		database.DropDBs(db)
	}
}
