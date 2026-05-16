package account

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindByEmail(email string) (*models.Account, error)
	FindByID(id uint) (*models.Account, error)
	Create(account *models.Account) error
	Update(account *models.Account) error
	GetAll() ([]models.Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewAccountrepository(db *gorm.DB) AccountRepository {
	return &repository{db: db}
}
