package year

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type YearRepository interface {
	Create(year *models.Year) error
}

type repository struct {
	db *gorm.DB
}

func NewYearRepository(db *gorm.DB) YearRepository {
	return &repository{db: db}
}
