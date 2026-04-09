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
	err := r.db.Where("cluster_id = ?", clusterId).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

// UpdatePollinationForm implements [PollinationRepository].
func (r *repository) UpdatePollinationForm(db *gorm.DB, form *models.PollinationForm) error {
	return commonRepo.Save(db, form)
}
