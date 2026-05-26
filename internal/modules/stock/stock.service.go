package stock

type StockService interface {
	CreateCarryOver(accountID uint, form CreateCarryOverRequest) (StockMovementResponse, error)
}
