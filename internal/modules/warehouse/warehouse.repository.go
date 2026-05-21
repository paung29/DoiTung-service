package warehouse

import "github.com/doitung/DoiTung-service/internal/models"

type WarehouseRepository interface {
	CreateNewWarehouse(form *models.Warehouse) error
	findAll() ([]models.Warehouse, error)
	findByName(warehouseName string) (*models.Warehouse, error)
	findById(warehouseId uint) (*models.Warehouse, error)
}
