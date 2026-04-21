package pollination

type PollinationFormRequest struct {
	ClusterId               uint   `json:"clusterId" validate:"required"`
	NumberPods              *uint  `json:"numberPods" validate:"required,number"`
	UnsuccessfulPollination *uint  `json:"unsuccessfulPollination" validate:"required,number"`
	Condition               string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PollinationFormResponse struct {
	Message string `json:"message"`
}

type PollinationFormDetails struct {
	ClusterId               uint   `json:"clusterId"`
	Location                string `json:"location"`
	PoleNo                  uint   `json:"poleNo"`
	ClusterNo               uint   `json:"clusterNo"`
	NumberPods              uint   `json:"numberPods"`
	UnsuccessfulPollination uint   `json:"unsuccessfulPollination"`
	GoodFlowers             uint   `json:"goodFlowers"`
	BadFlowers              uint   `json:"badFlowers"`
	Condition               string `json:"condition"`
	PollinationFormDone     bool   `json:"pollinationFormDone"`
}
