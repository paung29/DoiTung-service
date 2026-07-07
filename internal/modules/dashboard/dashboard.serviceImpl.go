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

func (s *service) GetConditionByStage(year int) (ConditionByStageResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(year)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ConditionByStageResponse{},
				utils.NotFoundError("year not found")
		}

		return ConditionByStageResponse{},
			utils.SystemError("failed to retrieve year")
	}

	yearId := yearRecord.YearID

	Stages := []struct {
		Name  string
		Table string
	}{{Name: "cluster", Table: "cluster_forms"},
		{Name: "flower", Table: "flower_forms"},
		{Name: "pollination", Table: "pollination_forms"},
		{Name: "pod", Table: "pod_forms"},
		{Name: "preHarvest", Table: "pre_harvest_forms"},
	}
	// get condition by stage of that year
	responseStages := make([]ConditionByStageItem, 0, len(Stages))

	for _, stage := range Stages {
		count, err := s.repo.GetConditionByStage(stage.Table, int(yearId))
		if err != nil {
			return ConditionByStageResponse{}, utils.SystemError("failed to retrieve condition by stage")
		}

		responseStages = append(responseStages, ConditionByStageItem{
			Stage:  stage.Name,
			Good:   count.Good,
			Insect: count.Insect,
			Rotten: count.Rotten,
		})

	}

	return ConditionByStageResponse{
		Year:   yearRecord.Year,
		Stages: responseStages,
	}, nil
}

func (s *service) GetFlowerProductionTrend() (FlowerProductionTrendResponse, error) {
	rows, err := s.repo.GetFlowerProductionTrend()
	if err != nil {
		return FlowerProductionTrendResponse{}, utils.SystemError("failed to retrieve flower production trend")
	}

	items := make([]FlowerProductionTrendItem, 0, len(rows))

	for _, row := range rows {
		items = append(items, FlowerProductionTrendItem{
			Year:         row.Year,
			TotalFlowers: row.TotalFlowers,
			GoodFlowers:  row.GoodFlowers,
			BadFlowers:   row.BadFlowers,
		})
	}

	return FlowerProductionTrendResponse{
		Items: items,
	}, nil
}

func (s *service) GetPodProductionTrend() (PodProductionTrendResponse, error) {
	rows, err := s.repo.GetPodOverviewTrend()
	if err != nil {
		return PodProductionTrendResponse{}, utils.SystemError("failed to retrieve pod production trend")
	}

	items := make([]PodProductionTrendItem, 0, len(rows))

	for _, row := range rows {
		items = append(items, PodProductionTrendItem{
			Year:          row.Year,
			TotalPods:     row.TotalPods,
			LostPods:      row.LostPods,
			RemainingPods: row.RemainingPods,
		})
	}

	return PodProductionTrendResponse{
		Items: items,
	}, nil
}

func (s *service) GetPodSetRateTrend() (PodSetRateTrendResponse, error) {
	rows, err := s.repo.GetPodSetRateTrend()
	if err != nil {
		return PodSetRateTrendResponse{}, utils.SystemError("failed to retrieve pod set rate trend")
	}

	items := make([]PodSetRateTrendItem, 0, len(rows))

	for _, row := range rows {
		totalFlowers := row.GoodFlowers + row.BadFlowers

		items = append(items, PodSetRateTrendItem{
			Year:                    row.Year,
			NumberPods:              row.NumberPods,
			UnsuccessfulPollination: row.UnsuccessfulPollination,
			GoodFlowers:             row.GoodFlowers,
			BadFlowers:              row.BadFlowers,
			TotalFlowers:            totalFlowers,
		})
	}

	return PodSetRateTrendResponse{
		Items: items,
	}, nil
}

func (s *service) GetHarvestablePodsTrend() (HarvestablePodsTrendResponse, error) {
	rows, err := s.repo.GetHarvestablePodsTrend()
	if err != nil {
		return HarvestablePodsTrendResponse{}, utils.SystemError("failed to retrieve harvestable pods trend")
	}

	items := make([]HarvestablePodsTrendItem, 0, len(rows))

	for _, row := range rows {
		items = append(items, HarvestablePodsTrendItem{
			Year:                  row.Year,
			TotalPods:             row.TotalPods,
			RemainingPods:         row.RemainingPods,
			SecondRoundPods:       row.SecondRoundPods,
			LostPodsBeforeHarvest: row.LostPodsBeforeHarvest,
			RemovedPods:           row.RemovedPods,
		})
	}

	return HarvestablePodsTrendResponse{
		Items: items,
	}, nil
}

func (s *service) GetFreshPodGradeTrend() (FreshPodGradeTrendResponse, error) {

	rows, err := s.repo.GetFreshPodGradeTrend()
	if err != nil {
		return FreshPodGradeTrendResponse{}, utils.SystemError("failed to retrieve fresh pod grade trend")
	}

	items := make([]FreshPodGradeTrendItem, 0, len(rows))

	for _, row := range rows {
		items = append(items, FreshPodGradeTrendItem{
			Year:       row.Year,
			GradeAPlus: row.GradeAPlus,
			GradeA:     row.GradeA,
			GradeB:     row.GradeB,
			GradeC:     row.GradeC,
			GradeDPlus: row.GradeDPlus,
			Undersized: row.Undersized,
			Rotten:     row.Rotten,
		})
	}

	return FreshPodGradeTrendResponse{
		Items: items,
	}, nil
}

func (s *service) GetProductivePolesTrend() (ProductivePolesTrendResponse, error) {
	rows, err := s.repo.GetProductivePolesTrend()
	if err != nil {
		return ProductivePolesTrendResponse{}, utils.SystemError("failed to retrieve productive poles trend")
	}

	items := make([]ProductivePolesTrendItem, 0, len(rows))

	for _, row := range rows {

		// Non-productive poles are poles which don't have harvest and grading forms.
		nonProductivePoles := row.TotalPoles - row.ProductivePoles
		if nonProductivePoles < 0 {
			nonProductivePoles = 0
		}
		items = append(items, ProductivePolesTrendItem{
			Year:               row.Year,
			TotalPoles:         row.TotalPoles,
			ProductivePoles:    row.ProductivePoles,
			NonProductivePoles: nonProductivePoles,
		})
	}

	return ProductivePolesTrendResponse{
		Items: items,
	}, nil
}

func (s *service) GetWeightPerPodTrend() (WeightPerPodTrendResponse, error) {
	rows, err := s.repo.GetWeightPerPodTrend()
	if err != nil {
		return WeightPerPodTrendResponse{}, utils.SystemError("failed to retrieve weight per pod trend")
	}

	items := make([]WeightPerPodTrendItem, 0, len(rows))

	for _, row := range rows {
		averageWeightPerPod := 0.0

		if row.TotalHarvestPods > 0 {
			averageWeightPerPod = row.TotalHarvestWeight / float64(row.TotalHarvestPods)
		}

		items = append(items, WeightPerPodTrendItem{
			Year:                row.Year,
			AverageWeightPerPod: utils.RoundTwoDecimals(averageWeightPerPod),
		})
	}

	return WeightPerPodTrendResponse{
		Items: items,
	}, nil
}
