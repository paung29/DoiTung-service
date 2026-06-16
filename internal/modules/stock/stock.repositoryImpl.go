package stock

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

type StockBalance struct {
	TotalGrams int
	TotalPods  int
}

type CustomerStockRow struct {
	CustomerId   int
	CustomerName string
	GradeA       int
	GradeB       int
	GradeC       int
	GradeFailed  int
	TotalWeight  int
	Note         *string
}

type GradeSummary struct {
	Grade      enums.Grade
	TotalGrams int
	TotalPods  int
}

type MonthlySummary struct {
	Month          int
	MonthName      string
	StockInWeight  int
	StockOutWeight int
	TotalWeight    int
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &repository{db: db}
}

func (r *repository) CreateStockMovement(db *gorm.DB, form *models.StockMovement) error {
	return commonrepo.Create(db, form)
}

func (r *repository) UpdateStockMovement(form *models.StockMovement) error {
	return commonrepo.Save(r.db, form)
}

func (r *repository) FindByID(id uint) (*models.StockMovement, error) {
	return commonrepo.FindByID[models.StockMovement](r.db, id)
}

func (r *repository) GetStockTotal(yearID uint, warehouseID uint, stockType enums.MovementType) (StockBalance, error) {
	var total StockBalance

	warehouseColumn := "to_warehouse_id"
	if stockType == enums.MovementIssued {
		warehouseColumn = "from_warehouse_id"
	}
	err := r.db.Model(&models.StockMovement{}).
		Select("COALESCE(SUM(total_grams), 0) as total_grams, COALESCE(SUM(total_pods), 0) as total_pods").
		Where("year_id = ? AND "+warehouseColumn+" = ? AND movement_type = ?", yearID, warehouseID, stockType).
		Scan(&total).Error
	return total, err
}

func (r *repository) GetStockBalanceForUpdate(db *gorm.DB, YearID uint, warehouseID uint, grade enums.Grade) (*models.StockBalance, error) {
	var balance models.StockBalance
	err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("year_id = ? AND warehouse_id = ? AND grade = ?", YearID, warehouseID, grade).
		First(&balance).Error
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (r *repository) CreateNewStockBalance(db *gorm.DB, form *models.StockBalance) error {
	return commonrepo.Create(db, form)
}

func (r *repository) UpdateStockBalance(db *gorm.DB, form *models.StockBalance) error {
	return db.Save(form).Error
}

func (r *repository) DeleteStockMovement(db *gorm.DB, id uint) error {
	return commonrepo.DeleteByID[models.StockMovement](db, id)
}

func (r *repository) GetAllByYearId(yearId uint) ([]*models.StockMovement, error) {
	var movements []*models.StockMovement
	err := r.db.
		Preload("ProductionYear").
		Preload("FromWarehouse").
		Preload("ToWarehouse").
		Where("year_id = ?", yearId).Find(&movements).Error
	return movements, err
}

func (r *repository) GetCustomerStockByYearId(yearId uint) ([]CustomerStockRow, error) {
	var rows []CustomerStockRow

	err := r.db.
		Table("stock_movements").
		Select(`
			stock_movements.issued_to_customer_id AS customer_id,
			customers.customer_name AS customer_name,

			COALESCE(SUM(CASE WHEN stock_movements.grade IN ('A_PLUS', 'A') THEN COALESCE(stock_movements.total_grams, 0) ELSE 0 END), 0) AS grade_a,
			COALESCE(SUM(CASE WHEN stock_movements.grade = 'B' THEN COALESCE(stock_movements.total_grams, 0) ELSE 0 END), 0) AS grade_b,
			COALESCE(SUM(CASE WHEN stock_movements.grade = 'C' THEN COALESCE(stock_movements.total_grams, 0) ELSE 0 END), 0) AS grade_c,
			COALESCE(SUM(CASE WHEN stock_movements.grade IN ('D', 'D_PLUS') THEN COALESCE(stock_movements.total_grams, 0) ELSE 0 END), 0) AS grade_failed,
			COALESCE(SUM(COALESCE(stock_movements.total_grams, 0)), 0) AS total_weight,

			customers.note AS note
		`).
		Joins("JOIN customers ON customers.customer_id = stock_movements.issued_to_customer_id").
		Where(
			"stock_movements.year_id = ? AND stock_movements.issued_to_customer_id IS NOT NULL AND stock_movements.movement_type = ?",
			yearId,
			enums.MovementIssued,
		).
		Group("stock_movements.issued_to_customer_id, customers.customer_name, customers.note").
		Scan(&rows).Error

	return rows, err
}

func (r *repository) GetStockOverviewBalanceByYearId(yearId uint) ([]GradeSummary, error) {
	var summaries []GradeSummary
	err := r.db.
		Model(&models.StockBalance{}).
		Select(`
	grade,
	COALESCE(SUM(total_grams), 0) AS total_grams,
	COALESCE(SUM(total_pods), 0) AS total_pods
        `).
		Where("year_id = ?", yearId).
		Group("grade").
		Scan(&summaries).Error
	return summaries, err
}

func (r *repository) GetIncomingStockTotal(yearId uint) (StockBalance, error) {
	var total StockBalance
	err := r.db.
		Model(&models.StockMovement{}).
		Select("COALESCE(SUM(total_grams), 0) as total_grams, COALESCE(SUM(total_pods), 0) as total_pods").
		Where("year_id = ? AND movement_type IN ?", yearId, []enums.MovementType{enums.MovementIncoming, enums.MovementCarryOver}).
		Scan(&total).Error
	return total, err
}

func (r *repository) GetIssuedStockTotal(yearId uint) (StockBalance, error) {
	var total StockBalance
	err := r.db.
		Model(&models.StockMovement{}).
		Select("COALESCE(SUM(total_grams), 0) as total_grams, COALESCE(SUM(total_pods), 0) as total_pods").
		Where("year_id = ? AND movement_type = ?", yearId, enums.MovementIssued).
		Scan(&total).Error
	return total, err
}

func (r *repository) GetMonthlySummary(yearId uint) ([]MonthlySummary, error) {
	var summaries []MonthlySummary
	err := r.db.
		Model(&models.StockMovement{}).
		Select(`
	EXTRACT(MONTH FROM recorded_date)::int AS month,
	COALESCE(SUM(CASE WHEN movement_type IN ('CARRY_OVER', 'INCOMING') THEN COALESCE(total_grams, 0) ELSE 0 END), 0) AS stock_in_weight,
	COALESCE(SUM(CASE WHEN movement_type = 'ISSUED' THEN COALESCE(total_grams, 0) ELSE 0 END), 0) AS stock_out_weight,
	COALESCE(SUM(CASE
		WHEN movement_type IN ('CARRY_OVER', 'INCOMING') THEN COALESCE(total_grams, 0)
		WHEN movement_type = 'ISSUED' THEN -COALESCE(total_grams, 0)
		ELSE 0
	END), 0) AS total_weight
`).
		Where("year_id = ?", yearId).
		Group("EXTRACT(MONTH FROM recorded_date)::int").
		Order("month ASC").Scan(&summaries).Error
	return summaries, err
}
