package pod

import "github.com/doitung/DoiTung-service/internal/modules/cluster"

type PodFormRequest struct {
	ClusterId uint   `json:"clusterId" validate:"required"`
	LostPods  *uint  `json:"lostPods" validate:"required,number"`
	Condition string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PodFormResponse struct {
	Message string `json:"message"`
}

type PodFormDetails struct {
	No            uint   `json:"no,omitempty"`
	ClusterId     uint   `json:"clusterId"`
	Location      string `json:"location"`
	PoleNo        uint   `json:"poleNo"`
	ClusterNo     uint   `json:"clusterNo"`
	NumberPods    uint   `json:"numberPods"`
	LostPods      uint   `json:"lostPods"`
	RemainingPods uint   `json:"remainingPods"`
	Condition     string `json:"condition"`
	PodFormDone   bool   `json:"podFormDone"`
	RecordedBy    string `json:"recordedBy,omitempty"`
	Date          string `json:"date,omitempty"`
}

type PodFormHistoriesResponse struct {
	PodFormHistories []cluster.ClusterInfo `json:"podFormHistories"`
}

type PodFormLists struct {
	PodForms []PodFormDetails `json:"podForms"`
}
