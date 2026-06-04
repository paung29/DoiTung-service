package pollination

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPollinationRepository(db *gorm.DB) PollinationRepository {
	return &repository{db: db}
}

// CreatePollinationForm implements [PollinationRepository].
func (r *repository) CreatePollinationForm(db *gorm.DB, form *models.PollinationForm) error {
	return commonRepo.Create(db, form)
}

// GetPollinationFormByClusterID implements [PollinationRepository].
func (r *repository) GetPollinationFormByClusterID(db *gorm.DB, clusterId uint) (*models.PollinationForm, error) {
	var form models.PollinationForm
	err := db.Where("cluster_id = ?", clusterId).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

// UpdatePollinationForm implements [PollinationRepository].
func (r *repository) UpdatePollinationForm(db *gorm.DB, form *models.PollinationForm) error {
	return commonRepo.Save(db, form)
}

// GetPollinationFormHistoriesByUserIdAndYearId implements [PollinationRepository].
func (r *repository) GetPollinationFormHistoriesByUserIdAndYearId(db *gorm.DB, userId uint, yearId uint) ([]models.PollinationForm, error) {
	var forms []models.PollinationForm
	err := db.Where("recorded_by_id = ? AND year_id = ?", userId, yearId).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}

func (r *repository) GetPollinationFormsByZoneId(db *gorm.DB, zoneId uint) ([]models.PollinationForm, error) {
	var forms []models.PollinationForm
	err := db.
		Model(&models.PollinationForm{}).
		Preload("RecordedBy").
		Preload("Cluster").
		Preload("Cluster.Pole").
		Preload("Cluster.Pole.Zone").
		Joins("JOIN clusters ON clusters.cluster_id = pollination_forms.cluster_id").
		Joins("JOIN poles ON poles.pole_id = clusters.pole_id").
		Joins("JOIN zones ON zones.zone_id = poles.zone_id").
		Where("zones.zone_id = ?", zoneId).
		Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}
