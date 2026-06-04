package zone

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type ZoneRepository interface {
	Create(db *gorm.DB, zone *models.Zone) error
	FindByYearAndZoneId(yearID uint, zoneId int) (*models.Zone, error)
	GetMaxZoneNoByYear(yearID uint) (int, error)
	FindByYearAndZoneName(yearID uint, name string) (*models.Zone, error)
	FindByYearID(yearID uint) ([]models.Zone, error)
	FindById(zoneId uint) (*models.Zone, error)
}
