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

func (r *repository) GetStockTotal(productionYearID uint, warehouseID uint, grade enums.Grade, stockType enums.MovementType) (StockBalance, error) {
	var total StockBalance

	warehouseColumn := "to_warehouse_id"
	if stockType == enums.MovementIssued {
		warehouseColumn = "from_warehouse_id"
	}
	err := r.db.Model(&models.StockMovement{}).
		Select("COALESCE(SUM(total_grams), 0) as total_grams, COALESCE(SUM(total_pods), 0) as total_pods").
		Where("production_year_id = ? AND "+warehouseColumn+" = ? AND grade = ? AND movement_type = ?", productionYearID, warehouseID, grade, stockType).
		Scan(&total).Error
	return total, err
}

func (r *repository) GetStockBalanceForUpdate(db *gorm.DB, productionYearID uint, warehouseID uint, grade enums.Grade) (*models.StockBalance, error) {
	var balance models.StockBalance
	err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("production_year_id = ? AND warehouse_id = ? AND grade = ?", productionYearID, warehouseID, grade).
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
