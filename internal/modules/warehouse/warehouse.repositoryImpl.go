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

func (r *repository) findAll() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	if err := r.db.Find(&warehouses).Error; err != nil {
		return nil, err
	}
	return warehouses, nil
}
