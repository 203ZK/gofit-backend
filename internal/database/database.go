package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/203ZK/gofit-backend/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB(cfg *config.DBConfig) error {
	dsn := os.ExpandEnv(cfg.Database.DSN)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("sqlOpen: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("connecting to database: %w", err)
	}

	fmt.Println("Connected to DB!")
	return nil
}
