package pole

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type PoleRepository interface {
	CreatePole(db *gorm.DB, pole *models.Pole) error
	GetPoleByZoneIdAndPoleNo(zoneId uint, poleNo uint) (*models.Pole, error)
	GetPoleById(poleId uint) (*models.Pole, error)
	UpdatePole(db *gorm.DB, pole *models.Pole) error
	UpdateHarvestGradingStatusByPoleId(poleId uint, status bool) error
	GetAllPolesByZoneId(zoneId uint) ([]models.Pole, error)
}
