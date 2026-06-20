package warehouse

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type StockBalance struct {
	TotalPods  uint
	TotalGrams float64
}

type WarehouseRepository interface {
	CreateNewWarehouse(form *models.Warehouse) error
	FindAll() ([]models.Warehouse, error)
	FindAllActive() ([]models.Warehouse, error)
	FindByName(warehouseName string) (*models.Warehouse, error)
	FindByID(warehouseId uint) (*models.Warehouse, error)
	UpdateWarehouse(warehouse *models.Warehouse) error
	GetStockTotal(YearID uint, warehouseID uint, stockType enums.MovementType) (StockBalance, error)
}
