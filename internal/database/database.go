package database

import (
	"fmt"
	"log"
	"os"

	"github.com/203ZK/gofit-backend/internal/config"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectToDB(cfg *config.DBConfig) (*gorm.DB, error) {
	dsn := os.ExpandEnv(cfg.Database.DSN)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	log.Println("Database connection successful")
	return DB, nil
}
