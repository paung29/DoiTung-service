package dashboard

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) SumTotalFlowers(yearId int) (int64, error) {
	var totalFlowers int64
	err := r.db.
		Model(&models.FlowerForm{}).
		Select("COALESCE(SUM(total_flowers), 0)").
		Where("year_id = ?", yearId).
		Scan(&totalFlowers).Error

	return totalFlowers, err
}

func (r *repository) GetPollinationStats(yearId int) (PollinationStats, error) {
	var stats PollinationStats
	err := r.db.
		Model(&models.PollinationForm{}).
		Select("COALESCE(SUM(good_flowers), 0) AS good_flowers, COALESCE(SUM(bad_flowers), 0) AS bad_flowers, COALESCE(SUM(number_pods), 0) AS total_pods").
		Where("year_id = ?", yearId).
		Scan(&stats).Error

	return stats, err
}

func (r *repository) GetPodStats(yearId int) (PodStats, error) {
	var stats PodStats

	err := r.db.
		Model(&models.PodForm{}).
		Select(`
			COALESCE(SUM(number_pods), 0) AS total_pods,
			COALESCE(SUM(lost_pods), 0) AS lost_pods,
			COALESCE(SUM(remaining_pods), 0) AS remaining_pods
		`).
		Where("year_id = ?", yearId).
		Scan(&stats).Error

	return stats, err
}

func (r *repository) GetHarvestStats(yearId int) (HarvestStats, error) {
	var stats HarvestStats

	err := r.db.
		Model(&models.HarvestGradingForm{}).
		Select(`
			COALESCE(SUM(
				grade_a_plus_weight +
				grade_a_weight +
				grade_b_weight +
				grade_c_weight +
				grade_d_plus_weight +
				undersized_weight +
				rotten_weight
			), 0) AS total_harvest_weight,

			COALESCE(SUM(
				grade_a_plus_count +
				grade_a_count +
				grade_b_count +
				grade_c_count +
				grade_d_plus_count +
				undersized_count +
				rotten_count
			), 0) AS total_harvest_pods
		`).
		Where("year_id = ?", yearId).
		Scan(&stats).Error

	return stats, err
}

func (r *repository) GetConditionByStage(tableName string, yearId int) (ConditionCountRow, error) {
	var row ConditionCountRow

	err := r.db.
		Table(tableName).
		Select(`
			COALESCE(SUM(CASE WHEN condition = 'GOOD' THEN 1 ELSE 0 END), 0) AS good,	
		COALESCE(SUM(CASE WHEN condition = 'INSECT' THEN 1 ELSE 0 END), 0) AS insect,
		COALESCE(SUM(CASE WHEN condition = 'ROTTEN' THEN 1 ELSE 0 END), 0) AS rotten
		`).
		Where("year_id = ?", yearId).
		Scan(&row).Error

	return row, err
}

func (r *repository) GetFlowerProductionTrend() ([]FlowerProductionTrendRow, error) {
	var rows []FlowerProductionTrendRow

	err := r.db.
		Table("years").
		Select(`
			years.year AS year,

			COALESCE((
				SELECT SUM(flower_forms.total_flowers)
				FROM flower_forms
				WHERE flower_forms.year_id = years.year_id
			), 0) AS total_flowers,

			COALESCE((
				SELECT SUM(pollination_forms.good_flowers)
				FROM pollination_forms
				WHERE pollination_forms.year_id = years.year_id
			), 0) AS good_flowers,

			COALESCE((
				SELECT SUM(pollination_forms.bad_flowers)
				FROM pollination_forms
				WHERE pollination_forms.year_id = years.year_id
			), 0) AS bad_flowers
		`).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}
