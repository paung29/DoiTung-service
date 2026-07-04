package dashboard

type DashboardRepository interface {
	SumTotalFlowers(yearId int) (int64, error)
	GetPollinationStats(yearId int) (PollinationStats, error)
	GetPodStats(yearId int) (PodStats, error)
	GetHarvestStats(yearId int) (HarvestStats, error)
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
