package year

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type YearRepository interface {
	Create(tx *gorm.DB, year *models.Year) error
	CreateFormSetting(tx *gorm.DB, setting *models.YearFormSetting) error
	FindByYear(year int) (*models.Year, error)
	FindByID(id uint) (*models.Year, error)
	FindFormSettingByYear(id uint) (*models.YearFormSetting, error)
	UpdateFormSetting(db *gorm.DB, setting *models.YearFormSetting) error
}


