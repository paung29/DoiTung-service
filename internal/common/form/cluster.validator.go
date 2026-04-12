package form

import (
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/utils"
)

type ClusterValidator struct {
	yearRepo    year.YearRepository
	zoneRepo    zone.ZoneRepository
	clusterRepo cluster.ClusterRepository
}

func NewClusterValidator(
	yearRepo year.YearRepository,
	zoneRepo zone.ZoneRepository,
	clusterRepo cluster.ClusterRepository,
) *ClusterValidator {
	return &ClusterValidator{
		yearRepo:    yearRepo,
		zoneRepo:    zoneRepo,
		clusterRepo: clusterRepo,
	}
}

func (v *ClusterValidator) ValidateClusterContext(
	year uint,
	zoneNo uint,
	poleNo uint,
	clusterNo uint,
	formType string,
) (uint, uint, error) {

	// Check if the year exists
	yearRecord, err := v.yearRepo.FindByYear(int(year))
	if err != nil {
		return 0, 0, utils.NotFoundError("year not found")
	}
	yearId := yearRecord.YearID

	// Check if the form setting is open for the year
	yearSetting, err := v.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return 0, 0, utils.NotFoundError("year setting not found")
	}

	// 3. Dynamic form check 🔥
	switch formType {
	case "flower":
		if !yearSetting.FlowerActive {
			return 0, 0, utils.BadRequestError("flower form is not open")
		}
	case "pollination":
		if !yearSetting.PollinationActive {
			return 0, 0, utils.BadRequestError("pollination form is not open")
		}
	case "cluster":
		if !yearSetting.ClusterActive {
			return 0, 0, utils.BadRequestError("cluster form is not open")
		}
	case "pod":
		if !yearSetting.PodActive {
			return 0, 0, utils.BadRequestError("pod form is not open")
		}
	case "preharvest":
		if !yearSetting.PreHarvestActive {
			return 0, 0, utils.BadRequestError("preharvest form is not open")
		}
	case "harvestgrading":
		if !yearSetting.HarvestGradingActive {
			return 0, 0, utils.BadRequestError("harvest grading form is not open")
		}
	}

	// Check if the zone exists
	zoneRecord, err := v.zoneRepo.FindByYearAndZoneNo(yearId, int(zoneNo))
	if err != nil {
		return 0, 0, utils.NotFoundError("zone not found")
	}

	// Check if the pole exists
	poleRecord, err := v.clusterRepo.FindPoleByZoneAndPoleNo(zoneRecord.ZoneID, poleNo)
	if err != nil {
		return 0, 0, utils.NotFoundError("pole not found")
	}

	// Check if the cluster exists
	clusterRecord, err := v.clusterRepo.FindClusterByPoleAndClusterNo(poleRecord.PoleID, clusterNo)
	if err != nil {
		return 0, 0, utils.NotFoundError("cluster not found")
	}

	return yearId, clusterRecord.ClusterID, nil
}
