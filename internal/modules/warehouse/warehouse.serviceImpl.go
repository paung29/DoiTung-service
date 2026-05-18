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

	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CreateWarehouseResponse{}, utils.NotFoundError("year not found")
		}
		return CreateWarehouseResponse{}, utils.SystemError("Failed to get Year")
	}

	yearId := yearRecord.YearID

	// Create the warehouse record
	warehouse := &models.Warehouse{
		YearID:        uint(yearId),
		WarehouseName: form.WarehouseName,
		ActiveStatus:  form.ActiveStatus,
	}

	if err := s.warehouseRepo.CreateNewWarehouse(warehouse); err != nil {
		if utils.IsDuplicateError(err) {
			return CreateWarehouseResponse{}, utils.ValidationError("Warehouse already exists", nil)
		}
		return CreateWarehouseResponse{}, utils.SystemError("Failed to create warehouse")
	}

	return CreateWarehouseResponse{
		Message: "Warehouse created successfully",
	}, nil
}
