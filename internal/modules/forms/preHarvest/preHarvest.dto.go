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

type PreHarvestFormDetails struct {
	ClusterId             uint   `json:"clusterId"`
	Location              string `json:"location"`
	PoleNo                uint   `json:"poleNo"`
	ClusterNo             uint   `json:"clusterNo"`
	RemainingPods         uint   `json:"remainingPods"`
	NumberPodsSecondRound uint   `json:"numberPodsSecondRound"`
	LostPodsBeforeHarvest uint   `json:"lostPodsBeforeHarvest"`
	RemovedPods           uint   `json:"removedPods"`
	PlantsRemoved         uint   `json:"plantsRemoved"`
	Condition             string `json:"condition"`
	PreHarvestFormDone    bool   `json:"preHarvestFormDone"`
}
