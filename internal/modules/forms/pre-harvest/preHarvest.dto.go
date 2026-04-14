package preharvest

type PreHarvestFormRequest struct {
	Year                  uint   `json:"year" validate:"required"`
	ZoneId                uint   `json:"zone-no" validate:"required"`
	PoleId                uint   `json:"pole-no" validate:"required"`
	ClusterId             uint   `json:"cluster-no" validate:"required"`
	NumberPodsSecondRound *uint  `json:"number-pods-second-round" validate:"required"`
	RemovedPods           *uint  `json:"removed-pods" validate:"required"`
	PlantsRemoved         *uint  `json:"plants-removed" validate:"required"`
	Condition             string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PreHarvestFormResponse struct {
	Message string `json:"message"`
}
