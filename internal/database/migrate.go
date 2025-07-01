package database

import "github.com/203ZK/gofit-backend/internal/models"

func MigrateDBs() {
	DB.AutoMigrate(&models.User{})
}
