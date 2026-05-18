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
