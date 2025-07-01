package database

import (
	"fmt"

	"github.com/203ZK/gofit-backend/internal/models"
	"gorm.io/gorm"
)

func SeedDBs(db *gorm.DB) error {
	users := []models.User{
		{Email: "test@gmail.com", EncryptedPassword: "setMeUp?"},
	}

	for _, u := range users {
		var user models.User
		result := db.Model(&models.User{}).Where("email = ?", u.Email).First(&user)
		if result.Error == gorm.ErrRecordNotFound {
			err := db.Create(&u).Error
			if err != nil {
				return fmt.Errorf("Error creating user: %w", err)
			}
		}
	}

	fmt.Println("Successfully seeded database!")
	return nil
}
