package year

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewYearRepository(db *gorm.DB) YearRepository {
	return &repository{db: db}
}

func (repo *repository) Create(tx *gorm.DB, year *models.Year) error {
	return commonrepo.Create(tx, year)
}

func (repo *repository) CreateFormSetting(tx *gorm.DB, setting *models.YearFormSetting) error {
	return commonrepo.Create(tx, setting)
}

func (repo *repository) FindByYear(yearValue int) (*models.Year, error) {
	var year models.Year
	if err := repo.db.Where("year = ?", yearValue).First(&year).Error; err != nil {
		return nil, err
	}
	return &year, nil
}

func (repo *repository) FindByID(id uint) (*models.Year, error) {
	return commonrepo.FindByID[models.Year](repo.db, id)
}

func (repo *repository) FindFormSettingByYear(id uint) (*models.YearFormSetting, error) {
	return commonrepo.FindByID[models.YearFormSetting](repo.db, id)
}

func (repo *repository) UpdateFormSetting(db *gorm.DB, setting *models.YearFormSetting) error {
	return commonrepo.Save(db, setting)
}

func (repo *repository) findAll() ([]models.Year, error) {
	var years []models.Year

	err := repo.db.Order("year asc").Find(&years).Error

	if err != nil {
		return nil, err
	}

	return years, nil
}

func (repo *repository) findAllYearDetails() ([]models.YearFormSetting, error) {
	var details []models.YearFormSetting
	err := repo.db.Preload("Year").Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}
