package preharvest

type PreHarvestFormRequest struct {
	Year                  uint   `json:"year" validate:"required"`
	ZoneId                uint   `json:"zoneNo" validate:"required"`
	PoleId                uint   `json:"poleNo" validate:"required"`
	ClusterId             uint   `json:"clusterNo" validate:"required"`
	NumberPodsSecondRound *uint  `json:"numberPodsSecondRound" validate:"required"`
	RemovedPods           *uint  `json:"removedPods" validate:"required"`
	PlantsRemoved         *uint  `json:"plantsRemoved" validate:"required"`
	Condition             string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PreHarvestFormResponse struct {
	Message string `json:"message"`
}
