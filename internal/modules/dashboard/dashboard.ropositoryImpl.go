package dashboard

import (
	"fmt"
	"strings"

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

func yearlySum(tableName string, columnName string, alias string) string {
	return fmt.Sprintf(`
		COALESCE((
			SELECT SUM(%s.%s)
			FROM %s
			WHERE %s.year_id = years.year_id
		), 0) AS %s
	`, tableName, columnName, tableName, tableName, alias)
}

func yearlySelect(parts ...string) string {
	return "years.year AS year,\n" + strings.Join(parts, ",\n")
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

	selectSQL := yearlySelect(
		yearlySum("flower_forms", "total_flowers", "total_flowers"),
		yearlySum("pollination_forms", "good_flowers", "good_flowers"),
		yearlySum("pollination_forms", "bad_flowers", "bad_flowers"),
	)

	err := r.db.
		Table("years").
		Select(selectSQL).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetPodOverviewTrend() ([]PodOverviewTrendRow, error) {
	var rows []PodOverviewTrendRow

	selectSQL := yearlySelect(
		yearlySum("pod_forms", "number_pods", "total_pods"),
		yearlySum("pod_forms", "lost_pods", "lost_pods"),
		yearlySum("pod_forms", "remaining_pods", "remaining_pods"),
	)

	err := r.db.
		Table("years").
		Select(selectSQL).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetPodSetRateTrend() ([]PodSetRateTrendRow, error) {
	var rows []PodSetRateTrendRow

	selectSQL := yearlySelect(
		yearlySum("pollination_forms", "number_pods", "number_pods"),
		yearlySum("pollination_forms", "unsuccessful_pollination", "unsuccessful_pollination"),
		yearlySum("pollination_forms", "good_flowers", "good_flowers"),
		yearlySum("pollination_forms", "bad_flowers", "bad_flowers"),
	)

	err := r.db.
		Table("years").
		Select(selectSQL).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetHarvestablePodsTrend() ([]HarvestablePodsTrendRow, error) {
	var rows []HarvestablePodsTrendRow

	selectSQL := yearlySelect(
		yearlySum("pod_forms", "number_pods", "total_pods"),
		yearlySum("pod_forms", "remaining_pods", "remaining_pods"),
		yearlySum("pre_harvest_forms", "number_pods_second_round", "second_round_pods"),
		yearlySum("pre_harvest_forms", "lost_pods_before_harvest", "lost_pods_before_harvest"),
		yearlySum("pre_harvest_forms", "removed_pods", "removed_pods"),
	)

	err := r.db.
		Table("years").
		Select(selectSQL).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetFreshPodGradeTrend() ([]FreshPodGradeTrendRow, error) {
	var rows []FreshPodGradeTrendRow

	selectSQL := yearlySelect(
		yearlySum("harvest_grading_forms", "grade_a_plus_weight", "grade_a_plus"),
		yearlySum("harvest_grading_forms", "grade_a_weight", "grade_a"),
		yearlySum("harvest_grading_forms", "grade_b_weight", "grade_b"),
		yearlySum("harvest_grading_forms", "grade_c_weight", "grade_c"),
		yearlySum("harvest_grading_forms", "grade_d_plus_weight", "grade_d_plus"),
		yearlySum("harvest_grading_forms", "undersized_weight", "undersized"),
		yearlySum("harvest_grading_forms", "rotten_weight", "rotten"),
	)

	err := r.db.
		Table("years").
		Select(selectSQL).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

// retrieve the weight and pod grater than 0 as productive poles
func (r *repository) GetProductivePolesTrend() ([]ProductivePolesTrendRow, error) {
	var rows []ProductivePolesTrendRow

	err := r.db.
		Table("years").
		Select(`
			years.year AS year,

			COALESCE((
				SELECT COUNT(poles.pole_id)
				FROM poles
				JOIN zones ON zones.zone_id = poles.zone_id
				WHERE zones.year_id = years.year_id
			), 0) AS total_poles,

			COALESCE((
			SELECT COUNT(DISTINCT harvest_grading_forms.pole_id)
			FROM harvest_grading_forms
			WHERE harvest_grading_forms.year_id = years.year_id
			AND (
			(	grade_a_plus_count +
				grade_a_count +
				grade_b_count +
				grade_c_count +
				grade_d_plus_count +
				undersized_count +
				rotten_count ) > 0
			OR
			(	grade_a_plus_weight +
				grade_a_weight +
				grade_b_weight +
				grade_c_weight +
				grade_d_plus_weight +
				undersized_weight +
				rotten_weight ) > 0 ) 
			), 0) AS productive_poles
		`).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetWeightPerPodTrend() ([]WeightPerPodTrendRow, error) {
	var rows []WeightPerPodTrendRow

	err := r.db.
		Table("years").
		Select(`
			years.year AS year,

			COALESCE((
				SELECT SUM(
					grade_a_plus_weight +
					grade_a_weight +
					grade_b_weight +
					grade_c_weight +
					grade_d_plus_weight +
					undersized_weight +
					rotten_weight
				)
				FROM harvest_grading_forms
				WHERE harvest_grading_forms.year_id = years.year_id
			), 0) AS total_harvest_weight,

			COALESCE((
				SELECT SUM(
					grade_a_plus_count +
					grade_a_count +
					grade_b_count +
					grade_c_count +
					grade_d_plus_count +
					undersized_count +
					rotten_count
				)
				FROM harvest_grading_forms
				WHERE harvest_grading_forms.year_id = years.year_id
			), 0) AS total_harvest_pods
		`).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetActualYieldTrend() ([]ActualYieldTrendRow, error) {
	var rows []ActualYieldTrendRow

	err := r.db.
		Table("years").
		Select(`
			years.year AS year,

			COALESCE((
				SELECT SUM(
					grade_a_plus_weight +
					grade_a_weight +
					grade_b_weight +
					grade_c_weight +
					grade_d_plus_weight +
					undersized_weight +
					rotten_weight
				)
				FROM harvest_grading_forms
				WHERE harvest_grading_forms.year_id = years.year_id
			), 0) AS total_harvest_weight,

			COALESCE((
				SELECT COUNT(DISTINCT harvest_grading_forms.pole_id)
				FROM harvest_grading_forms
				WHERE harvest_grading_forms.year_id = years.year_id
				AND (
					(
						grade_a_plus_count +
						grade_a_count +
						grade_b_count +
						grade_c_count +
						grade_d_plus_count +
						undersized_count +
						rotten_count
					) > 0
					OR
					(
						grade_a_plus_weight +
						grade_a_weight +
						grade_b_weight +
						grade_c_weight +
						grade_d_plus_weight +
						undersized_weight +
						rotten_weight
					) > 0
				)
			), 0) AS productive_poles
		`).
		Order("years.year ASC").
		Scan(&rows).Error

	return rows, err
}
