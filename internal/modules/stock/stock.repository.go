package stock

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateStockMovement(db *gorm.DB, form *models.StockMovement) error
	UpdateStockMovement(form *models.StockMovement) error
	FindByID(id uint) (*models.StockMovement, error)
	CreateNewStockBalance(db *gorm.DB, form *models.StockBalance) error
	GetStockBalanceForUpdate(db *gorm.DB, productionYearID uint, warehouseID uint, grade enums.Grade) (*models.StockBalance, error)
	UpdateStockBalance(db *gorm.DB, form *models.StockBalance) error
	GetStockTotal(productionYearID uint, warehouseID uint, grade enums.Grade, stockType enums.MovementType) (StockBalance, error)
	DeleteStockMovement(db *gorm.DB, id uint) error
}
