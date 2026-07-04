package dashboard

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	yearRepo year.YearRepository
	repo     DashboardRepository
}

func NewDashboardService(yearRepo year.YearRepository, repo DashboardRepository) DashboardService {
	return &service{
		yearRepo: yearRepo,
		repo:     repo,
	}
}

func (s *service) GetPerformanceOverview(year int) (PerformanceOverviewResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(year)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PerformanceOverviewResponse{},
				utils.NotFoundError("year not found")
		}

		return PerformanceOverviewResponse{},
			utils.SystemError("failed to retrieve year")
	}

	yearId := yearRecord.YearID

	// get total flowers of that year
	totalFlowers, err := s.repo.SumTotalFlowers(int(yearId))
	if err != nil {
		return PerformanceOverviewResponse{}, utils.SystemError("failed to retrieve total flowers")
	}

	// get pollination stats of that year
	pollinationStats, err := s.repo.GetPollinationStats(int(yearId))
	if err != nil {
		return PerformanceOverviewResponse{}, utils.SystemError("failed to retrieve pollination stats")
	}

	// get pod stats of that year
	podStats, err := s.repo.GetPodStats(int(yearId))
	if err != nil {
		return PerformanceOverviewResponse{}, utils.SystemError("failed to retrieve pod stats")
	}

	// get harvest stats of that year
	harvestStats, err := s.repo.GetHarvestStats(int(yearId))
	if err != nil {
		return PerformanceOverviewResponse{}, utils.SystemError("failed to retrieve harvest stats")
	}

	// calculate flower loss rate and pod success rate (by dividing bad flowers by total flowers (good flowers + bad flowers))
	flowerLossRate := utils.CalculateRate(
		pollinationStats.BadFlowers,
		pollinationStats.GoodFlowers+pollinationStats.BadFlowers,
	)

	// calculate pod success rate (by dividing good flowers by total pods)
	podSuccessRate := utils.CalculateRate(
		pollinationStats.TotalPods,
		pollinationStats.GoodFlowers,
	)

	return PerformanceOverviewResponse{
		Year:               yearRecord.Year,
		TotalFlowers:       totalFlowers,
		TotalPods:          podStats.TotalPods,
		FlowerLossRate:     flowerLossRate,
		PodSuccessRate:     podSuccessRate,
		TotalHarvestWeight: utils.RoundTwoDecimals(harvestStats.TotalHarvestWeight),
		TotalHarvestPods:   harvestStats.TotalHarvestPods,
	}, nil
}
