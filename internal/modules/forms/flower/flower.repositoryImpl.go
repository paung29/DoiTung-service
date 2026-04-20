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

// GetFlowerFormDetailsByClusterID implements [FlowerRepository].
func (r *repository) GetFlowerFormDetailsByClusterID(db *gorm.DB, clusterId uint) (*models.FlowerForm, error) {
	var form models.FlowerForm
	if error := r.db.Preload("Cluster").Preload("Cluster.Pole").Preload("Cluster.Pole.Zone").Where("cluster_id = ?", clusterId).First(&form).Error; error != nil {
		return nil, error
	}
	return &form, nil
}
