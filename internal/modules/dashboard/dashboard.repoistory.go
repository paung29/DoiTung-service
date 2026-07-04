package dashboard

type DashboardRepository interface {
	SumTotalFlowers(yearId int) (int64, error)
	GetPollinationStats(yearId int) (PollinationStats, error)
	GetPodStats(yearId int) (PodStats, error)
	GetHarvestStats(yearId int) (HarvestStats, error)
	GetConditionByStage(tableName string, yearId int) (ConditionCountRow, error)
	GetFlowerProductionTrend() ([]FlowerProductionTrendRow, error)
}

type PollinationStats struct {
	GoodFlowers int64
	BadFlowers  int64
	TotalPods   int64
}

type PodStats struct {
	TotalPods     int64
	LostPods      int64
	RemainingPods int64
}

type HarvestStats struct {
	TotalHarvestWeight float64
	TotalHarvestPods   int64
}

type ConditionCountRow struct {
	Good   int64
	Insect int64
	Rotten int64
}

type FlowerProductionTrendRow struct {
	Year         int
	TotalFlowers int64
	GoodFlowers  int64
	BadFlowers   int64
}
