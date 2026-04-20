package flower

type FlowerFormRequest struct {
	Year         uint   `json:"year" validate:"required,number"`
	ZoneNo       uint   `json:"zoneNo" validate:"required,number"`
	PoleNo       uint   `json:"poleNo" validate:"required,number"`
	ClusterNo    uint   `json:"clusterNo" validate:"required,number"`
	TotalFlowers *uint  `json:"totalFlowers" validate:"required,number"`
	Condition    string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type FlowerFormResponse struct {
	Message string `json:"message"`
}

type FlowerFormDetails struct {
	ClusterId      uint   `json:"clusterId"`
	Location       string `json:"location"`
	PoleNo         int    `json:"poleNo"`
	ClusterNo      int    `json:"clusterNo"`
	TotalFlowers   uint   `json:"totalFlowers"`
	Condition      string `json:"condition"`
	FlowerFormDone bool   `json:"flowerFormDone"`
}
