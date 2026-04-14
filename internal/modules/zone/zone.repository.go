package zone

import (
	"gorm.io/gorm"
	"github.com/doitung/DoiTung-service/internal/models"
)

type ZoneRepository interface {
	Create(db *gorm.DB, zone *models.Zone) error
	FindByYearAndZoneNo(yearID uint, zoneNo int) (*models.Zone, error)
	GetMaxZoneNoByYear(yearID uint) (int, error)
	FindByYearAndZoneName(yearID uint, name string) (*models.Zone, error)
	FindByYearID(yearID uint) ([]models.Zone, error)
}
