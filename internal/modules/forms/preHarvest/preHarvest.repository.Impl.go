package preharvest

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPreHarvestRepository(db *gorm.DB) PreHarvestRepository {
	return &repository{db: db}
}

// CreateOrUpdatePreHarvestForm implements [PreHarvestRepository].
func (r *repository) CreateOrUpdatePreHarvestForm(db *gorm.DB, form *models.PreHarvestForm) error {
	return commonRepo.Create(db, form)
}

// GetPreHarvestFormByClusterId implements [PreHarvestRepository].
func (r *repository) GetPreHarvestFormByClusterId(db *gorm.DB, clusterId uint) (*models.PreHarvestForm, error) {
	var form models.PreHarvestForm
	err := r.db.Where("cluster_id = ?", clusterId).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

// UpdatePreHarvestForm implements [PreHarvestRepository].
func (r *repository) UpdatePreHarvestForm(db *gorm.DB, form *models.PreHarvestForm) error {
	return commonRepo.Save(db, form)
}
