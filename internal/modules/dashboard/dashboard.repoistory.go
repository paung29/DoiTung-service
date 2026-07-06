package dashboard

type DashboardRepository interface {
	SumTotalFlowers(yearId int) (int64, error)
	GetPollinationStats(yearId int) (PollinationStats, error)
	GetPodStats(yearId int) (PodStats, error)
	GetHarvestStats(yearId int) (HarvestStats, error)
	GetConditionByStage(tableName string, yearId int) (ConditionCountRow, error)
	GetFlowerProductionTrend() ([]FlowerProductionTrendRow, error)
	GetPodOverviewTrend() ([]PodOverviewTrendRow, error)
	GetPodSetRateTrend() ([]PodSetRateTrendRow, error)
	GetHarvestablePodsTrend() ([]HarvestablePodsTrendRow, error)
	GetFreshPodGradeTrend() ([]FreshPodGradeTrendRow, error)
	GetProductivePolesTrend() ([]ProductivePolesTrendRow, error)
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

type PodOverviewTrendRow struct {
	Year          int
	TotalPods     int64
	LostPods      int64
	RemainingPods int64
}

type PodSetRateTrendRow struct {
	Year                    int
	NumberPods              int64
	UnsuccessfulPollination int64
	GoodFlowers             int64
	BadFlowers              int64
}

type HarvestablePodsTrendRow struct {
	Year                  int
	TotalPods             int64
	RemainingPods         int64
	SecondRoundPods       int64
	LostPodsBeforeHarvest int64
	RemovedPods           int64
}

type FreshPodGradeTrendRow struct {
	Year       int
	GradeAPlus float64
	GradeA     float64
	GradeB     float64
	GradeC     float64
	GradeDPlus float64
	Undersized float64
	Rotten     float64
}

type ProductivePolesTrendRow struct {
	Year            int
	TotalPoles      int64
	ProductivePoles int64
}
