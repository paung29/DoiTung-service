package warehouse

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db            *gorm.DB
	yearRepo      year.YearRepository
	warehouseRepo WarehouseRepository
}

func NewWarehouseService(db *gorm.DB, yearRepo year.YearRepository, warehouseRepo WarehouseRepository) WarehouseService {
	return &service{
		db:            db,
		yearRepo:      yearRepo,
		warehouseRepo: warehouseRepo,
	}
}

func (s *service) CreateWarehouse(form CreateWarehouseRequest) (CreateWarehouseResponse, error) {

	warehouseRecord, err := s.warehouseRepo.FindByName(form.WarehouseName)

	if err == nil && warehouseRecord != nil {
		return CreateWarehouseResponse{}, utils.ValidationError("Warehouse already exists", nil)
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return CreateWarehouseResponse{}, utils.SystemError("Failed to check existing warehouse")
	}

	// Create the warehouse record
	warehouse := &models.Warehouse{
		WarehouseName: form.WarehouseName,
		ActiveStatus:  form.ActiveStatus,
	}

	if err := s.warehouseRepo.CreateNewWarehouse(warehouse); err != nil {
		return CreateWarehouseResponse{}, utils.SystemError("Failed to create warehouse")
	}

	return CreateWarehouseResponse{
		Message: "Warehouse created successfully",
	}, nil
}

func (s *service) GetAllWarehouses() (GetAllWarehousesResponse, error) {
	warehouses, err := s.warehouseRepo.FindAllActive()
	if err != nil {
		return GetAllWarehousesResponse{}, utils.SystemError("Failed to retrieve warehouses")
	}

	warehouseDetails := make([]WarehouseDetail, 0, len(warehouses))
	for _, warehouse := range warehouses {
		warehouseDetails = append(warehouseDetails, WarehouseDetail{
			WarehouseId:   warehouse.WarehouseID,
			WarehouseName: warehouse.WarehouseName,
			ActiveStatus:  warehouse.ActiveStatus,
		})
	}

	return GetAllWarehousesResponse{
		Warehouses: warehouseDetails,
	}, nil
}

func (s *service) GetWarehouseById(warehouseId uint) (WarehouseDetail, error) {
	warehouse, err := s.warehouseRepo.FindByID(warehouseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return WarehouseDetail{}, utils.ValidationError("Warehouse not found", nil)
		}
		return WarehouseDetail{}, utils.SystemError("Failed to retrieve warehouse")
	}

	return WarehouseDetail{
		WarehouseId:   warehouse.WarehouseID,
		WarehouseName: warehouse.WarehouseName,
		ActiveStatus:  warehouse.ActiveStatus,
	}, nil
}

func (s *service) UpdateWarehouse(form UpdateWarehouseRequest) (UpdateWarehouseResponse, error) {
	warehouse, err := s.warehouseRepo.FindByID(form.WarehouseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UpdateWarehouseResponse{}, utils.ValidationError("Warehouse not found", nil)
		}
		return UpdateWarehouseResponse{}, utils.SystemError("Failed to retrieve warehouse")
	}

	warehouseRecord, err := s.warehouseRepo.FindByName(form.WarehouseName)

	if err == nil && warehouseRecord != nil && warehouseRecord.WarehouseID != form.WarehouseId {
		return UpdateWarehouseResponse{}, utils.ValidationError("Warehouse name already exists", nil)
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return UpdateWarehouseResponse{}, utils.SystemError("Failed to check existing warehouse")
	}

	warehouse.WarehouseName = form.WarehouseName
	warehouse.ActiveStatus = form.ActiveStatus

	if err := s.warehouseRepo.UpdateWarehouse(warehouse); err != nil {
		return UpdateWarehouseResponse{}, utils.SystemError("Failed to update warehouse")
	}

	return UpdateWarehouseResponse{
		Message: "Warehouse updated successfully",
	}, nil
}

func (s *service) GetWarehouseTableByYear(year int) (WarehouseTableByYearResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(year)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return WarehouseTableByYearResponse{}, utils.ValidationError("Year not found", nil)
		}
		return WarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve year")
	}

	warehouses, err := s.warehouseRepo.FindAll()
	if err != nil {
		return WarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve warehouses")
	}

	var warehouseTableYearResponse WarehouseTableByYearResponse
	warehouseTableYearResponse.TotalWarehouses = len(warehouses)

	activeWarehouses, err := s.warehouseRepo.FindAllActive()
	if err != nil {
		return WarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve active warehouses")
	}
	warehouseTableYearResponse.TotalActiveWarehouses = len(activeWarehouses)

	totalStocksPods := 0
	totalStocksGrams := 0.0
	warehouseTable := make([]WarehouseTableItem, 0, len(warehouses))
	for _, warehouse := range warehouses {

		totalIncoming, err := s.warehouseRepo.GetStockTotal(yearRecord.YearID, warehouse.WarehouseID, enums.MovementIncoming)
		if err != nil {
			return WarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve incoming stock balance")
		}
		totalCarryOver, err := s.warehouseRepo.GetStockTotal(yearRecord.YearID, warehouse.WarehouseID, enums.MovementCarryOver)
		if err != nil {
			return WarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve carry-over stock balance")
		}

		totalPods := totalIncoming.TotalPods + totalCarryOver.TotalPods
		totalWeights := totalIncoming.TotalGrams + totalCarryOver.TotalGrams

		totalIssued, err := s.warehouseRepo.GetStockTotal(yearRecord.YearID, warehouse.WarehouseID, enums.MovementIssued)
		if err != nil {
			return WarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve issued stock balance")
		}
		totalDistributedPods := totalIssued.TotalPods
		totalDistributedWeights := totalIssued.TotalGrams

		remainingPods := totalPods - totalDistributedPods
		remainingWeights := totalWeights - totalDistributedWeights

		warehouseTable = append(warehouseTable, WarehouseTableItem{
			WarehouseId:   warehouse.WarehouseID,
			WarehouseName: warehouse.WarehouseName,
			ActiveStatus:  warehouse.ActiveStatus,

			TotalPods:    int(totalPods),
			TotalWeights: totalWeights,

			DistributedPods:    int(totalDistributedPods),
			DistributedWeights: totalDistributedWeights,

			RemainingPods:    int(remainingPods),
			RemainingWeights: remainingWeights,
		})
		totalStocksPods += int(remainingPods)
		totalStocksGrams += remainingWeights
	}
	warehouseTableYearResponse.WarehouseTable = warehouseTable
	warehouseTableYearResponse.TotalStocksPods = totalStocksPods
	warehouseTableYearResponse.TotalStocksWeights = totalStocksGrams

	return warehouseTableYearResponse, nil
}
