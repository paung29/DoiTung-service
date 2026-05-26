package stock

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &repository{db: db}
}

func (r *repository) CreateStockMovement(form *models.StockMovement) error {
	return commonrepo.Create(r.db, form)
}
