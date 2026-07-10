package zone

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db       *gorm.DB
	zoneRepo ZoneRepository
	yearRepo year.YearRepository
}

func NewZoneService(db *gorm.DB, zoneRepo ZoneRepository, yearRepo year.YearRepository) ZoneService {
	return &service{
		db:       db,
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
		YearID:   yearID,
		ZoneName: form.Name,
		ZoneNo:   nextZoneNo,
	}

	if err := s.zoneRepo.Create(s.db, zone); err != nil {
		return CreateZoneResponse{}, utils.SystemError("failed to create zone")
	}

	return CreateZoneResponse{
		Message: "zone created successfully",
	}, nil
}

func (s service) GetAllZone(yearID uint) (GetAllZoneResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(int(yearID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return GetAllZoneResponse{}, utils.NotFoundError("year not found")
		}
		return GetAllZoneResponse{}, utils.SystemError("failed to check year")
	}

	zones, err := s.zoneRepo.FindByYearID(yearRecord.YearID)
	if err != nil {
		return GetAllZoneResponse{}, utils.SystemError("failed to get zones")
	}

	zoneResponses := make([]ZoneResponse, 0, len(zones))
	for _, z := range zones {
		zoneResponses = append(zoneResponses, ZoneResponse{
			ZoneID:   z.ZoneID,
			ZoneName: z.ZoneName,
		})
	}

	return GetAllZoneResponse{
		Zones: zoneResponses,
	}, nil
}

func (s service) GetZoneManagementTable(yearID uint) (GetZoneManagementTableResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(int(yearID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return GetZoneManagementTableResponse{}, utils.NotFoundError("year not found")
		}
		return GetZoneManagementTableResponse{}, utils.SystemError("failed to check year")
	}

	zones, err := s.zoneRepo.FindByYearID(yearRecord.YearID)
	if err != nil {
		return GetZoneManagementTableResponse{}, utils.SystemError("failed to get zones")
	}

	var zoneInfos []ZoneManagementInfo
	var totalPoles int64 = 0

	for _, z := range zones {
		polesInZone, err := s.zoneRepo.GetTotalPolesByZoneId(z.ZoneID)
		if err != nil {
			return GetZoneManagementTableResponse{}, utils.SystemError("failed to count poles in zone")
		}

		totalPoles += polesInZone

		zoneInfos = append(zoneInfos, ZoneManagementInfo{
			ZoneID:           z.ZoneID,
			ZoneName:         z.ZoneName,
			TotalPolesInZone: polesInZone,
		})
	}

	return GetZoneManagementTableResponse{
		TotalZones: len(zones),
		TotalPoles: totalPoles,
		Zones:      zoneInfos,
	}, nil
}

func (s service) UpdateZoneName(form UpdateZoneName) (UpdateZoneNameResponse, error) {
	zone, err := s.zoneRepo.FindById(form.ZoneID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UpdateZoneNameResponse{}, utils.NotFoundError("zone not found")
		}
	}

	newZoneName := form.ZoneName

	existingZone, err := s.zoneRepo.FindByYearAndZoneName(zone.YearID, newZoneName)
	if err == nil && existingZone != nil && existingZone.ZoneID != zone.ZoneID {
		return UpdateZoneNameResponse{}, utils.BadRequestError("zone name already exists in this year")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return UpdateZoneNameResponse{}, utils.SystemError("failed to check zone name")
	}

	zone.ZoneName = newZoneName

	if err := s.zoneRepo.UpdateZoneName(s.db, zone.ZoneID, newZoneName); err != nil {
		return UpdateZoneNameResponse{}, utils.SystemError("failed to update zone name")
	}

	return UpdateZoneNameResponse{
		Message: "zone name updated successfully",
	}, nil
}
