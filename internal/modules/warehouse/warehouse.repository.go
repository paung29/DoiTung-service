package warehouse

import "github.com/doitung/DoiTung-service/internal/models"

type WarehouseRepository interface {
	CreateNewWarehouse(form *models.Warehouse) error
	FindAll() ([]models.Warehouse, error)
	FindByName(warehouseName string) (*models.Warehouse, error)
	FindByID(warehouseId uint) (*models.Warehouse, error)
	UpdateWarehouse(warehouse *models.Warehouse) error
}
