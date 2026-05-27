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
	repo          StockRepository
	yearRepo      year.YearRepository
	warehouseRepo warehouse.WarehouseRepository
	customerRepo  customer.CustomerRepository
}

func NewStockService(repo StockRepository, yearRepo year.YearRepository, warehouseRepo warehouse.WarehouseRepository, customerRepo customer.CustomerRepository) StockService {
	return &service{repo: repo, yearRepo: yearRepo, warehouseRepo: warehouseRepo, customerRepo: customerRepo}
}

func (s *service) CreateCarryOver(accountID uint, form CreateCarryOverStockRequest) (StockMovementResponse, error) {

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
		MovementType:     enums.MovementCarryOver,
	}
	err = s.repo.CreateStockMovement(stockMovement)
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
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
	err = s.repo.CreateStockMovement(stockMovement)
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
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

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}

	customerRecord, err := s.customerRepo.FindByCustomerID(*form.CustomerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Customer doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve customer record")
	}

	stockMovement := &models.StockMovement{
		RecordedByID:       accountID,
		YearID:             yearRecord.YearID,
		ProductionYearID:   &ProductYearRecord.YearID,
		FromWarehouseID:    &warehouseRecord.WarehouseID,
		IssuedToCustomerID: &customerRecord.CustomerID,
		Grade:              form.Grade,
		PricePerGram:       form.PricePerGram,
		TotalGrams:         form.TotalGrams,
		TotalPods:          form.TotalPods,
		Details:            form.Details,
		RecordedDate:       form.RecordedDate,
		MovementType:       enums.MovementIssued,
	}
	err = s.repo.CreateStockMovement(stockMovement)
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
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
