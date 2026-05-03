package pod

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type PodRepository interface {
	CreatePodForm(db *gorm.DB, form *models.PodForm) error
	GetPodFormByClusterId(db *gorm.DB, clusterId uint) (*models.PodForm, error)
	UpdatePodForm(db *gorm.DB, form *models.PodForm) error
	GetPodFormHistoriesByUserId(db *gorm.DB, userId uint) ([]models.PodForm, error)
}
