package stock

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type StockRepository interface {
	CreateStockMovement(form *models.StockMovement) error
	UpdateStockMovement(form *models.StockMovement) error
	FindByID(id uint) (*models.StockMovement, error)
	GetStockTotal(productionYearID uint, warehouseID uint, grade enums.Grade, stockType enums.MovementType) (StockBalance, error)
}
