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

type ConditionByStageResponse struct {
	Year   int                    `json:"year"`
	Stages []ConditionByStageItem `json:"stages"`
}

type ConditionByStageItem struct {
	Stage  string `json:"stage"`
	Good   int64  `json:"good"`
	Insect int64  `json:"insect"`
	Rotten int64  `json:"rotten"`
}

type FlowerProductionTrendResponse struct {
	Items []FlowerProductionTrendItem `json:"items"`
}

type FlowerProductionTrendItem struct {
	Year         int   `json:"year"`
	TotalFlowers int64 `json:"totalFlowers"`
	GoodFlowers  int64 `json:"goodFlowers"`
	BadFlowers   int64 `json:"badFlowers"`
}

type PodProductionTrendResponse struct {
	Items []PodProductionTrendItem `json:"items"`
}

type PodProductionTrendItem struct {
	Year          int   `json:"year"`
	TotalPods     int64 `json:"totalPods"`
	LostPods      int64 `json:"lostPods"`
	RemainingPods int64 `json:"remainingPods"`
}
