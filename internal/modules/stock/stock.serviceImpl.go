package stock

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
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
}

func NewStockService(repo StockRepository, yearRepo year.YearRepository, warehouseRepo warehouse.WarehouseRepository) StockService {
	return &service{repo: repo, yearRepo: yearRepo, warehouseRepo: warehouseRepo}
}

func (s *service) CreateCarryOver(accountID uint, form CreateCarryOverRequest) (StockMovementResponse, error) {

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
