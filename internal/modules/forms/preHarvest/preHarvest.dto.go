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
	No                    uint   `json:"no,omitempty"`
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
	RecordedBy            string `json:"recordedBy,omitempty"`
	Date                  string `json:"date,omitempty"`
}

type PreHarvestFormHistoriesResponse struct {
	PreHarvestForms []PreHarvestFormHistory `json:"preHarvestForms"`
}

type PreHarvestFormHistory struct {
	No           uint   `json:"no"`
	ClusterId    uint   `json:"clusterId"`
	Location     string `json:"location"`
	PoleNo       uint   `json:"poleNo"`
	ClusterNo    uint   `json:"clusterNo"`
	ProgressDone uint   `json:"progressDone"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type PreHarvestFormLists struct {
	PreHarvestForms []PreHarvestFormDetails `json:"preHarvestForms"`
}
