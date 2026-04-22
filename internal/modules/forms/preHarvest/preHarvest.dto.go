package preharvest

type PreHarvestFormRequest struct {
	ClusterId             uint   `json:"clusterId" validate:"required"`
	NumberPodsSecondRound *uint  `json:"numberPodsSecondRound" validate:"required"`
	RemovedPods           *uint  `json:"removedPods" validate:"required"`
	PlantsRemoved         *uint  `json:"plantsRemoved" validate:"required"`
	Condition             string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PreHarvestFormResponse struct {
	Message string `json:"message"`
}
