package warehouse

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"

	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &repository{db: db}
}

func (r *repository) CreateNewWarehouse(form *models.Warehouse) error {
	return commonrepo.Create(r.db, form)
}

func (r *repository) FindAll() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	if err := r.db.Where("active_status = ?", true).Order("warehouse_id ASC").Find(&warehouses).Error; err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (r *repository) FindByName(warehouseName string) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	if err := r.db.Where("warehouse_name = ?", warehouseName).First(&warehouse).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *repository) FindByID(warehouseId uint) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	if err := r.db.Where("warehouse_id = ?", warehouseId).First(&warehouse).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *repository) UpdateWarehouse(warehouse *models.Warehouse) error {
	return commonrepo.Save(r.db, warehouse)
}
