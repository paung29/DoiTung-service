package flower

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type FlowerRepository interface {
	CreateFlowerForm(db *gorm.DB, form *models.FlowerForm) error
	GetFlowerFormByClusterID(db *gorm.DB, clusterId uint) (*models.FlowerForm, error)
	UpdateFlowerForm(db *gorm.DB, form *models.FlowerForm) error
	GetFlowerFormDetailsByClusterID(db *gorm.DB, clusterId uint) (*models.FlowerForm, error)
	GetFlowerFormHistoriesByUserIdAndYearId(db *gorm.DB, userId uint, yearId uint) ([]models.FlowerForm, error)
	GetFlowerFormsByZoneId(db *gorm.DB, zoneId uint) ([]models.FlowerForm, error)
}
