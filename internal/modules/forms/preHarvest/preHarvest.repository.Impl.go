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

// GetPreHarvestFormsByUserIdAndYear implements [PreHarvestRepository].
func (r *repository) GetPreHarvestFormsByUserIdAndYear(db *gorm.DB, userId uint, year uint) ([]models.PreHarvestForm, error) {
	var forms []models.PreHarvestForm
	err := db.Where("recorded_by_id = ? AND year_id = ?", userId, year).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}

func (r *repository) GetPreHarvestFormsByZoneId(db *gorm.DB, zoneId uint) ([]models.PreHarvestForm, error) {
	var forms []models.PreHarvestForm
	err := db.
		Model(&models.PreHarvestForm{}).
		Preload("RecordedBy").
		Preload("Cluster").
		Preload("Cluster.Pole").
			Preload("Cluster.Pole.Zone").
			Joins("JOIN clusters ON clusters.cluster_id = pre_harvest_forms.cluster_id").
			Joins("JOIN poles ON poles.pole_id = clusters.pole_id").
			Where("poles.zone_id = ?", zoneId).
			Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}
