package harvestgrading

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type HarvestGradingRepository interface {
	CreateHarvestGradingForm(db *gorm.DB, form *models.HarvestGradingForm) error
	GetHarvestGradingFormByPoleId(poleId uint) (*models.HarvestGradingForm, error)
	UpdateHarvestGradingForm(form *models.HarvestGradingForm) error
	GetHarvestGradingFormsByUserIdAndYearId(db *gorm.DB, userId uint, yearId uint) ([]models.HarvestGradingForm, error)
	GetHarvestGradingFormsByZoneId(db *gorm.DB, zoneId uint) ([]models.HarvestGradingForm, error)
}
