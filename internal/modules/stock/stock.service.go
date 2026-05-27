package stock

type StockService interface {
	CreateCarryOver(accountID uint, form CreateCarryOverRequest) (StockMovementResponse, error)
	CreateIncomingStock(accountID uint, form CreateIncomingStockRequest) (StockMovementResponse, error)
	CreateIssuedStock(accountID uint, form CreateIssuedStockRequest) (StockMovementResponse, error)
}
