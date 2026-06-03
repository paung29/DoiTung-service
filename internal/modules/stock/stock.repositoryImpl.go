package stock

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"gorm.io/gorm"
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

func (r *repository) CreateStockMovement(form *models.StockMovement) error {
	return commonrepo.Create(r.db, form)
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
