package warehouse

import "github.com/doitung/DoiTung-service/internal/models"

type WarehouseRepository interface {
	CreateNewWarehouse(form *models.Warehouse) error
}
