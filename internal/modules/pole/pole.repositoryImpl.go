package pole

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPoleRepository(db *gorm.DB) PoleRepository {
	return &repository{db: db}
}

// CreatePole implements [PoleRepository].
func (r *repository) CreatePole(db *gorm.DB, pole *models.Pole) error {
	return commonRepo.Create(db, pole)
}

// GetPoleByPoleNo implements [PoleRepository].
func (r *repository) GetPoleByZoneIdAndPoleNo(zoneId uint, poleNo uint) (*models.Pole, error) {
	var pole models.Pole
	err := r.db.Where(" zone_id = ? AND pole_no = ?", zoneId, poleNo).First(&pole).Error
	if err != nil {
		return nil, err
	}
	return &pole, nil
}

// UpdateHarvestGradingStatusByPoleId implements [PoleRepository].
func (r *repository) UpdateHarvestGradingStatusByPoleId(poleId uint, status bool) error {
	return r.db.Model(&models.Pole{}).Where("pole_id = ?", poleId).Update("harvest_grading_form_done", status).Error
}

// UpdatePole implements [PoleRepository].
func (r *repository) UpdatePole(db *gorm.DB, pole *models.Pole) error {
	return commonRepo.Save(db, pole)
}

func (r *repository) GetAllPolesByZoneId(zoneId uint) ([]models.Pole, error) {
	var poles []models.Pole
	err := r.db.Where("zone_id = ?", zoneId).Find(&poles).Error
	if err != nil {
		return nil, err
	}
	return poles, nil
}
