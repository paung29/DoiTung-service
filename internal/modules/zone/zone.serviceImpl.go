package zone

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	zoneRepo ZoneRepository
	yearRepo year.YearRepository
}

func NewZoneService(db *gorm.DB, zoneRepo ZoneRepository, yearRepo year.YearRepository) ZoneService{
	return &service{
		db: db,
		zoneRepo: zoneRepo,
		yearRepo: yearRepo,
	}
}

func (s service) CreateZone(form CreateZoneRequest) (CreateZoneResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CreateZoneResponse{}, utils.NotFoundError("year not found")
		}
		return CreateZoneResponse{}, utils.SystemError("failed to check year")
	}

	yearID := yearRecord.YearID

	existingZone, err := s.zoneRepo.FindByYearAndZoneName(yearID, form.Name)

	if err == nil && existingZone != nil {
		return CreateZoneResponse{}, utils.BadRequestError("zone name already exists in this year")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return CreateZoneResponse{}, utils.SystemError("failed to check zone name")
	}

	maxZoneNo, err := s.zoneRepo.GetMaxZoneNoByYear(yearID)
	if err != nil {
		return CreateZoneResponse{}, utils.SystemError("failed to generate zone number")
	}

	nextZoneNo := maxZoneNo + 1

	zone := &models.Zone{
		YearID: yearID,
		ZoneName: form.Name,
		ZoneNo: nextZoneNo,
	}

	if err := s.zoneRepo.Create(s.db, zone); err != nil {
		return CreateZoneResponse{}, utils.SystemError("failed to create zone")
	}

	return CreateZoneResponse{
		Message: "zone created successfully",
	}, nil
}