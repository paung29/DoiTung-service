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

	warehouseRecord, err := s.warehouseRepo.findByName(form.WarehouseName)

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
	warehouses, err := s.warehouseRepo.findAll()
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
