package dashboard

type PerformanceOverviewResponse struct {
	Year               int     `json:"year"`
	TotalFlowers       int64   `json:"totalFlowers"`
	TotalPods          int64   `json:"totalPods"`
	FlowerLossRate     float64 `json:"flowerLossRate"`
	PodSuccessRate     float64 `json:"podSuccessRate"`
	TotalHarvestWeight float64 `json:"totalHarvestWeight"`
	TotalHarvestPods   int64   `json:"totalHarvestPods"`
}
