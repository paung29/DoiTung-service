package zone

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
	commonrepo"github.com/doitung/DoiTung-service/internal/common/repository"
)

type repository struct {
	db *gorm.DB
}

func NewZoneRepository(db *gorm.DB) ZoneRepository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Create(db *gorm.DB, zone *models.Zone) error {
	return commonrepo.Create(db, zone)
}

func (repo *repository) FindByYearAndZoneNo(yearID uint, zoneNo int) (*models.Zone, error) {
	var zone models.Zone

	if err := repo.db.Where("year_id = ? AND zone_no = ?", yearID, zoneNo).First(&zone).Error; err != nil {
		return nil, err
	}

	return &zone, nil
}

func (repo *repository) GetMaxZoneNoByYear(yearID uint) (int, error) {
	var max int

	err := repo.db.Model(&models.Zone{}).
					Where("year_id = ?", yearID).
					Select("COALESCE(MAX(zone_no), 0)").
					Scan(&max).Error

	if err != nil {
		return 0, nil
	}

	return max, nil
}
