package exportdata

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewExportDataRepository(db *gorm.DB) ExportDataRepository {
	return &repository{db: db}
}

func (r *repository) ExportClusterFormsXLSX(yearID uint) ([]models.Cluster, error) {
	var clusters []models.Cluster
	err := r.db.
		Joins("JOIN poles ON poles.pole_id = clusters.pole_id").
		Joins("JOIN zones ON zones.zone_id = poles.zone_id").
		Preload("Pole").
		Preload("Pole.Zone").
		Preload("ClusterForms", "year_id = ?", yearID).
		Preload("FlowerForms", "year_id = ?", yearID).
		Preload("PollinationForms", "year_id = ?", yearID).
		Preload("PodForms", "year_id = ?", yearID).
		Preload("PreHarvestForms", "year_id = ?", yearID).
		Where("zones.year_id = ?", yearID).
		Order("zones.zone_no ASC, poles.pole_no ASC, clusters.cluster_no ASC").
		Find(&clusters).Error

	return clusters, err
}

func (r *repository) FindHarvestGradingFormsByYearID(
	yearID uint,
) ([]models.HarvestGradingForm, error) {
	var forms []models.HarvestGradingForm

	err := r.db.
		Joins(
			"JOIN poles ON poles.pole_id = harvest_grading_forms.pole_id",
		).
		Joins(
			"JOIN zones ON zones.zone_id = poles.zone_id",
		).
		Preload("Pole").
		Preload("Pole.Zone").
		Where("harvest_grading_forms.year_id = ?", yearID).
		Order("zones.zone_no ASC, poles.pole_no ASC").
		Find(&forms).Error

	return forms, err
}

func (r *repository) FindStockMovementsByYearID(yearID uint) ([]models.StockMovement, error) {
	var movements []models.StockMovement

	err := r.db.
		Preload("ProductionYear").
		Preload("FromWarehouse").
		Preload("ToWarehouse").
		Preload("IssuedToCustomer").
		Where("year_id = ?", yearID).
		Order("recorded_date ASC, stock_movement_id ASC").
		Find(&movements).Error

	return movements, err
}
