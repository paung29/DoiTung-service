package pollination

type PollinationFormRequest struct {
	Year                    uint   `json:"year" validate:"required,number"`
	ZoneNo                  uint   `json:"zone-no" validate:"required,number"`
	PoleNo                  uint   `json:"pole-no" validate:"required,number"`
	ClusterNo               uint   `json:"cluster-no" validate:"required,number"`
	NumberPods              *uint  `json:"number-pods" validate:"required,number"`
	UnsuccessfulPollination *uint  `json:"unsuccessful-pollination" validate:"required,number"`
	Condition               string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PollinationFormResponse struct {
	Message string `json:"message"`
}
