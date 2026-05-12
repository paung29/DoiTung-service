package preharvest

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type PreHarvestRepository interface {
	CreateOrUpdatePreHarvestForm(db *gorm.DB, form *models.PreHarvestForm) error
	GetPreHarvestFormByClusterId(db *gorm.DB, clusterId uint) (*models.PreHarvestForm, error)
	UpdatePreHarvestForm(db *gorm.DB, form *models.PreHarvestForm) error
	GetPreHarvestFormsByUserIdAndYear(db *gorm.DB, userId uint, year uint) ([]models.PreHarvestForm, error)
}
