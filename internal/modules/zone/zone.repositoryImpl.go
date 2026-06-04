package zone

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
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

func (repo *repository) FindByYearAndZoneId(yearID uint, zoneId int) (*models.Zone, error) {
	var zone models.Zone

	if err := repo.db.Where("year_id = ? AND zone_id = ?", yearID, zoneId).First(&zone).Error; err != nil {
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

func (repo *repository) FindByYearAndZoneName(yearID uint, name string) (*models.Zone, error) {
	var zone models.Zone

	if err := repo.db.Where("year_id = ? AND zone_name = ?", yearID, name).First(&zone).Error; err != nil {
		return nil, err
	}

	return &zone, nil
}

func (repo *repository) FindByYearID(yearID uint) ([]models.Zone, error) {
	var zones []models.Zone
	err := repo.db.Where("year_id = ?", yearID).Find(&zones).Error

	return zones, err
}

func (repo *repository) FindById(zoneId uint) (*models.Zone, error) {
	return commonrepo.FindByID[models.Zone](repo.db, zoneId)
}
