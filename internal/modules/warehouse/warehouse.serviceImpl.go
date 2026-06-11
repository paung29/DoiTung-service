package warehouse

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/year"
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

	var warehouseDetails []WarehouseDetail
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

// func (s *service) GetWarehouseTableByYear(year int) (GetWarehouseTableByYearResponse, error) {

// 	yearRecord, err := s.yearRepo.FindByYear(year)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return GetWarehouseTableByYearResponse{}, utils.ValidationError("Year not found", nil)
// 		}
// 		return GetWarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve year")
// 	}

// 	warehouseTable, err := s.warehouseRepo.GetWarehouseTableByYear(yearRecord.YearID)
// 	if err != nil {
// 		return GetWarehouseTableByYearResponse{}, utils.SystemError("Failed to retrieve warehouse table")
// 	}

// 	return warehouseTable, nil
// }
