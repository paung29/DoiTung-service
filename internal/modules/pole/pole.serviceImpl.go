package pole

import (
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type poleService struct {
	db       *gorm.DB
	yearRepo year.YearRepository
	zoneRepo zone.ZoneRepository
	poleRepo PoleRepository
}

func NewPoleService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, poleRepo PoleRepository) PoleService {
	return &poleService{
		db:       db,
		yearRepo: yearRepo,
		zoneRepo: zoneRepo,
		poleRepo: poleRepo,
	}
}

func (s *poleService) GetPoleByZone(year int, zoneNo int) (PolesByZoneResponse, error) {
	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(year)
	if err != nil {
		return PolesByZoneResponse{}, utils.NotFoundError("year not found")
	}
	yearId := yearRecord.YearID

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneId(uint(yearId), zoneNo)
	if err != nil {
		return PolesByZoneResponse{}, utils.NotFoundError("zone not found")
	}
	zoneId := zoneRecord.ZoneID

	poles, err := s.poleRepo.GetAllPolesByZoneId(zoneId)
	if err != nil {
		return PolesByZoneResponse{}, utils.SystemError("failed to get poles by zone")
	}
	var poleResponses []PoleResponse
	for _, pole := range poles {
		poleResponses = append(poleResponses, PoleResponse{
			PoleId:                 pole.PoleID,
			ZoneId:                 pole.ZoneID,
			Location:               pole.Zone.ZoneName,
			PoleNo:                 uint(pole.PoleNo),
			HarvestGradingFormDone: pole.HarvestGradingFormDone,
			CreatedAt:              pole.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:              pole.UpdatedAt.Format("2006-01-02 15:04"),
		})
	}
	return PolesByZoneResponse{Poles: poleResponses}, nil
}

func (s *poleService) GetPoleFilter(zoneId uint, poleNo *uint, harvestGradingFormDone *bool) (PoleFilterResponse, error) {
	zoneRecord, err := s.zoneRepo.FindById(zoneId)
	if err != nil {
		return PoleFilterResponse{}, utils.NotFoundError("zone not found")
	}

	poles, err := s.poleRepo.GetPolesByFilter(zoneRecord.ZoneID, poleNo, harvestGradingFormDone)
	if err != nil {
		return PoleFilterResponse{}, utils.SystemError("failed to get poles by filter")
	}

	poleResponses := make([]PoleResponse, 0, len(poles))

	for _, pole := range poles {
		poleResponses = append(poleResponses, PoleResponse{
			PoleId:                 pole.PoleID,
			ZoneId:                 pole.ZoneID,
			Location:               pole.Zone.ZoneName,
			PoleNo:                 uint(pole.PoleNo),
			HarvestGradingFormDone: pole.HarvestGradingFormDone,
			CreatedAt:              pole.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:              pole.UpdatedAt.Format("2006-01-02 15:04"),
		})
	}

	return PoleFilterResponse{
		Poles: poleResponses,
	}, nil
}
