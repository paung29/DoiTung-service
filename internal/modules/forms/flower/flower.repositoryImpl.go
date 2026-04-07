package flower

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewFlowerRepository(db *gorm.DB) FlowerRepository {
	return &repository{db: db}
}

// GetFlowerFormByClusterID implements [FlowerRepository].
func (r *repository) GetFlowerFormByClusterID(db *gorm.DB, clusterId uint) (*models.FlowerForm, error) {
	var form models.FlowerForm
	err := r.db.Where("cluster_id = ?", clusterId).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

// UpdateFlowerForm implements [FlowerRepository].
func (r *repository) UpdateFlowerForm(db *gorm.DB, form *models.FlowerForm) error {
	return commonRepo.Save(db, form)
}

// CreateFlowerForm implements [FlowerRepository].
func (r *repository) CreateFlowerForm(db *gorm.DB, form *models.FlowerForm) error {
	return commonRepo.Create(db, form)
}
