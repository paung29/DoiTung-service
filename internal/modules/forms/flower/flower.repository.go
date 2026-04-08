package flower

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type FlowerRepository interface {
	CreateFlowerForm(db *gorm.DB, form *models.FlowerForm) error
	GetFlowerFormByClusterID(db *gorm.DB, clusterId uint) (*models.FlowerForm, error)
	UpdateFlowerForm(db *gorm.DB, form *models.FlowerForm) error
}
