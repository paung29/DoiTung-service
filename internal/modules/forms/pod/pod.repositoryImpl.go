package pod

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPodRepository(db *gorm.DB) PodRepository {
	return &repository{db: db}
}

// CreatePodForm implements [PodRepository].
func (r *repository) CreatePodForm(db *gorm.DB, form *models.PodForm) error {
	return commonRepo.Create(db, form)
}

// GetPodFormByClusterId implements [PodRepository].
func (r *repository) GetPodFormByClusterId(db *gorm.DB, clusterId uint) (*models.PodForm, error) {
	var form models.PodForm
	err := r.db.Where("cluster_id = ?", clusterId).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

// UpdatePodForm implements [PodRepository].
func (r *repository) UpdatePodForm(db *gorm.DB, form *models.PodForm) error {
	return commonRepo.Save(db, form)
}

// GetPodFormHistoriesByUserId implements [PodRepository].
func (r *repository) GetPodFormHistoriesByUserId(db *gorm.DB, userId uint) ([]models.PodForm, error) {
	var forms []models.PodForm
	err := db.Where("recorded_by_id = ?", userId).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}
