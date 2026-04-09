package pollination

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type PollinationRepository interface {
	CreatePollinationForm(db *gorm.DB, form *models.PollinationForm) error
	GetPollinationFormByClusterID(db *gorm.DB, clusterId uint) (*models.PollinationForm, error)
	UpdatePollinationForm(db *gorm.DB, form *models.PollinationForm) error
}
