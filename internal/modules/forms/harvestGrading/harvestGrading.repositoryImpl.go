package harvestgrading

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewHarvestGradingRepository(db *gorm.DB) HarvestGradingRepository {
	return &repository{db: db}
}

// CreateHarvestGradingForm implements [HarvestGradingRepository].
func (r *repository) CreateHarvestGradingForm(db *gorm.DB, form *models.HarvestGradingForm) error {
	return commonRepo.Create(db, form)
}

// GetHarvestGradingFormByPoleId implements [HarvestGradingRepository].
func (r *repository) GetHarvestGradingFormByPoleId(poleId uint) (*models.HarvestGradingForm, error) {
	var form models.HarvestGradingForm
	err := r.db.Where("pole_id = ?", poleId).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

// UpdateHarvestGradingForm implements [HarvestGradingRepository].
func (r *repository) UpdateHarvestGradingForm(form *models.HarvestGradingForm) error {
	return commonRepo.Save(r.db, form)
}

// GetHarvestGradingFormsByUserIdAndYear implements [HarvestGradingRepository].
func (r *repository) GetHarvestGradingFormsByUserIdAndYearId(db *gorm.DB, userId uint, yearId uint) ([]models.HarvestGradingForm, error) {
	var forms []models.HarvestGradingForm
	err := db.Where("recorded_by_id = ? AND year_id = ?", userId, yearId).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}
