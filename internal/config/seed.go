package config

import(
	"log"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"golang.org/x/crypto/bcrypt"
)

func SeedAccounts() {

	var count int64

	DB.Model(&models.Account{}).Where("email = ?", "admin@doitung.com").Count(&count)

	if count > 0 {
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash password")
	}

	admin := models.Account{
		Name:         "Admin",
		Email:        "admin@doitung.com",
		PasswordHash: string(password),
		Role:         enums.RoleAdmin,
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Fatal("failed to seed admin account")
	}

	log.Println("Default admin account created")
}