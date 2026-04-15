package pollination

type PollinationFormRequest struct {
	Year                    uint   `json:"year" validate:"required,number"`
	ZoneNo                  uint   `json:"zoneNo" validate:"required,number"`
	PoleNo                  uint   `json:"poleNo" validate:"required,number"`
	ClusterNo               uint   `json:"clusterNo" validate:"required,number"`
	NumberPods              *uint  `json:"numberPods" validate:"required,number"`
	UnsuccessfulPollination *uint  `json:"unsuccessfulPollination" validate:"required,number"`
	Condition               string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PollinationFormResponse struct {
	Message string `json:"message"`
}
