package stock

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/customer"
	"github.com/doitung/DoiTung-service/internal/modules/warehouse"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db            *gorm.DB
	repo          StockRepository
	yearRepo      year.YearRepository
	warehouseRepo warehouse.WarehouseRepository
	customerRepo  customer.CustomerRepository
}

func NewStockService(db *gorm.DB, repo StockRepository, yearRepo year.YearRepository, warehouseRepo warehouse.WarehouseRepository, customerRepo customer.CustomerRepository) StockService {
	return &service{db: db, repo: repo, yearRepo: yearRepo, warehouseRepo: warehouseRepo, customerRepo: customerRepo}
}

func (s *service) CreateCarryOver(accountID uint, form CreateCarryOverStockRequest) (StockMovementResponse, error) {

	yearRecord, err := s.yearRepo.FindByID(form.YearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve year record")
	}
	yearId := yearRecord.YearID

	ProductYearRecord, err := s.yearRepo.FindByID(*form.ProductionYearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
	}
	productionYearId := ProductYearRecord.YearID

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}
	warehouseId := warehouseRecord.WarehouseID

	if form.TotalGrams == nil && form.TotalPods == nil {
		return StockMovementResponse{}, utils.BadRequestError("Total grams or total pods must be provided")
	}

	// transaction start here
	tx := s.db.Begin()
	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}

	// Create stock movement record if fail rollback transaction and return error
	stockMovement := &models.StockMovement{
		RecordedByID:     accountID,
		YearID:           yearId,
		ProductionYearID: &productionYearId,
		ToWarehouseID:    &warehouseId,
		Grade:            form.Grade,
		TotalGrams:       form.TotalGrams,
		TotalPods:        form.TotalPods,
		Details:          form.Details,
		RecordedDate:     form.RecordedDate,
		MovementType:     enums.MovementCarryOver,
	}
	err = s.repo.CreateStockMovement(tx, stockMovement)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
	}

	stockBalanceRecord, err := s.repo.GetStockBalanceForUpdate(tx, productionYearId, warehouseId, form.Grade)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no existing stock balance record, create a new one with the carry over quantity
			stockBalanceRecord = &models.StockBalance{
				ProductionYearID: productionYearId,
				WarehouseID:      warehouseId,
				Grade:            form.Grade,
				TotalGrams:       0,
				TotalPods:        0,
			}
			err = s.repo.CreateNewStockBalance(tx, stockBalanceRecord)
			if err != nil {
				tx.Rollback()
				return StockMovementResponse{}, utils.SystemError("Failed to create new stock balance")
			}
		} else {
			tx.Rollback()
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
		}
	}

	// Update stock balance by adding the carry over quantity to the existing balance
	if form.TotalGrams != nil {
		stockBalanceRecord.TotalGrams += *form.TotalGrams
	}
	if form.TotalPods != nil {
		stockBalanceRecord.TotalPods += *form.TotalPods
	}
	err = s.repo.UpdateStockBalance(tx, stockBalanceRecord)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}
	if err = tx.Commit().Error; err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement created successfully"}, nil
}

func (s *service) CreateIncomingStock(accountID uint, form CreateIncomingStockRequest) (StockMovementResponse, error) {

	yearRecord, err := s.yearRepo.FindByID(form.YearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve year record")
	}

	ProductYearRecord, err := s.yearRepo.FindByID(*form.ProductionYearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
	}

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}

	if form.TotalGrams == nil && form.TotalPods == nil {
		return StockMovementResponse{}, utils.BadRequestError("Total grams or total pods must be provided")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}
	stockMovement := &models.StockMovement{
		RecordedByID:     accountID,
		YearID:           yearRecord.YearID,
		ProductionYearID: &ProductYearRecord.YearID,
		ToWarehouseID:    &warehouseRecord.WarehouseID,
		Grade:            form.Grade,
		TotalGrams:       form.TotalGrams,
		TotalPods:        form.TotalPods,
		Details:          form.Details,
		RecordedDate:     form.RecordedDate,
		MovementType:     enums.MovementIncoming,
	}
	err = s.repo.CreateStockMovement(tx, stockMovement)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
	}

	stockBalanceRecord, err := s.repo.GetStockBalanceForUpdate(tx, ProductYearRecord.YearID, warehouseRecord.WarehouseID, form.Grade)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stockBalanceRecord = &models.StockBalance{
				ProductionYearID: ProductYearRecord.YearID,
				WarehouseID:      warehouseRecord.WarehouseID,
				Grade:            form.Grade,
				TotalGrams:       0,
				TotalPods:        0,
			}
			err = s.repo.CreateNewStockBalance(tx, stockBalanceRecord)
			if err != nil {
				tx.Rollback()
				return StockMovementResponse{}, utils.SystemError("Failed to create new stock balance")
			}
		} else {
			tx.Rollback()
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
		}
	}

	if form.TotalGrams != nil {
		stockBalanceRecord.TotalGrams += *form.TotalGrams
	}
	if form.TotalPods != nil {
		stockBalanceRecord.TotalPods += *form.TotalPods
	}
	err = s.repo.UpdateStockBalance(tx, stockBalanceRecord)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}
	if err = tx.Commit().Error; err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement created successfully"}, nil
}

func (s *service) CreateIssuedStock(accountID uint, form CreateIssuedStockRequest) (StockMovementResponse, error) {

	yearRecord, err := s.yearRepo.FindByID(form.YearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve year record")
	}

	ProductYearRecord, err := s.yearRepo.FindByID(*form.ProductionYearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
	}
	productionYearId := ProductYearRecord.YearID

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}
	warehouseId := warehouseRecord.WarehouseID

	customerRecord, err := s.customerRepo.FindByCustomerID(*form.CustomerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Customer doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve customer record")
	}

	if int(form.TotalGrams) <= 0 && int(form.TotalPods) <= 0 {
		return StockMovementResponse{}, utils.BadRequestError("Total grams or total pods must be greater than 0")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}
	// Get total stock for the specified production year, warehouse, and grade
	stockBalance, err := s.repo.GetStockBalanceForUpdate(tx, productionYearId, warehouseId, form.Grade)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
	}

	if stockBalance == nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("No stock available for the specified production year, warehouse, and grade")
	}

	if stockBalance.TotalGrams == 0 && stockBalance.TotalPods == 0 {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("No stock available for the specified production year, warehouse, and grade")
	}

	// If the total stock is less than the requested issued quantity, return an error
	if stockBalance.TotalGrams < int(form.TotalGrams) || stockBalance.TotalPods < int(form.TotalPods) {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("Insufficient stock available for the requested issue quantity")
	}

	stockMovement := &models.StockMovement{
		RecordedByID:       accountID,
		YearID:             yearRecord.YearID,
		ProductionYearID:   &productionYearId,
		FromWarehouseID:    &warehouseId,
		IssuedToCustomerID: &customerRecord.CustomerID,
		Grade:              form.Grade,
		PricePerGram:       &form.PricePerGram,
		TotalGrams:         &form.TotalGrams,
		TotalPods:          &form.TotalPods,
		Details:            form.Details,
		RecordedDate:       form.RecordedDate,
		MovementType:       enums.MovementIssued,
	}
	err = s.repo.CreateStockMovement(tx, stockMovement)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
	}

	// Update stock balance by subtracting the issued quantity from the existing balance
	stockBalance.TotalGrams -= int(form.TotalGrams)
	stockBalance.TotalPods -= int(form.TotalPods)
	err = s.repo.UpdateStockBalance(tx, stockBalance)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}

	err = tx.Commit().Error
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement created successfully"}, nil
}

func (s *service) UpdateStockMovement(accountID uint, form UpdateStockMovementRequest) (StockMovementResponse, error) {

	stockMovementRecord, err := s.repo.FindByID(form.StockMovementID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Stock movement record doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock movement record")
	}

	if form.ProductionYearID != nil {
		ProductYearRecord, err := s.yearRepo.FindByID(*form.ProductionYearID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
			}
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
		}
		stockMovementRecord.ProductionYearID = &ProductYearRecord.YearID
	}

	if form.WarehouseID != nil {
		warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
			}
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
		}
		if stockMovementRecord.MovementType == enums.MovementIssued {
			stockMovementRecord.FromWarehouseID = &warehouseRecord.WarehouseID
		} else {
			stockMovementRecord.ToWarehouseID = &warehouseRecord.WarehouseID
		}
	}

	// Check if the movement type is issued before validating customer ID and price per gram, as these fields are only relevant for issued stock movements
	isMomentIssued := stockMovementRecord.MovementType == enums.MovementIssued
	if form.CustomerID != nil && isMomentIssued {
		customerRecord, err := s.customerRepo.FindByCustomerID(*form.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return StockMovementResponse{}, utils.BadRequestError("Customer doesn't exist")
			}
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve customer record")
		}
		stockMovementRecord.IssuedToCustomerID = &customerRecord.CustomerID
	}

	stockMovementRecord.Grade = form.Grade
	if form.PricePerGram != nil && isMomentIssued {
		stockMovementRecord.PricePerGram = form.PricePerGram
	}
	if form.TotalGrams != nil {
		stockMovementRecord.TotalGrams = form.TotalGrams
	}
	if form.TotalPods != nil {
		stockMovementRecord.TotalPods = form.TotalPods
	}
	if form.Details != nil {
		stockMovementRecord.Details = form.Details
	}
	stockMovementRecord.RecordedByID = accountID

	err = s.repo.UpdateStockMovement(stockMovementRecord)
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to update stock movement")
	}

	return StockMovementResponse{Message: "Stock movement updated successfully"}, nil
}
