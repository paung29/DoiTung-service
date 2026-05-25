package stock

import "github.com/doitung/DoiTung-service/internal/models"

type StockRepository interface {
	CreateStockMovement(form *models.StockMovement) error
}
