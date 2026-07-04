package dashboard

type DashboardService interface {
	GetPerformanceOverview(year int) (PerformanceOverviewResponse, error)
	GetConditionByStage(year int) (ConditionByStageResponse, error)
	GetFlowerProductionTrend() (FlowerProductionTrendResponse, error)
	GetPodProductionTrend() (PodProductionTrendResponse, error)
	GetPodSetRateTrend() (PodSetRateTrendResponse, error)
	GetHarvestablePodsTrend() (HarvestablePodsTrendResponse, error)
}
