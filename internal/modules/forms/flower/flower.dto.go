package flower

type FlowerFormRequest struct {
	ClusterId    uint   `json:"clusterId" validate:"required"`
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
