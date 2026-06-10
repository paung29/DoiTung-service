package stock

type StockService interface {
	CreateCarryOver(accountID uint, form CreateCarryOverStockRequest) (StockMovementResponse, error)
	CreateIncomingStock(accountID uint, form CreateIncomingStockRequest) (StockMovementResponse, error)
	CreateIssuedStock(accountID uint, form CreateIssuedStockRequest) (StockMovementResponse, error)
	// UpdateStockMovement(accountID uint, form UpdateStockMovementRequest) (StockMovementResponse, error)
	DeleteStockMovement(stockMovementID uint) (StockMovementResponse, error)
}
