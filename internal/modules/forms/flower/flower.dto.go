package flower

type FlowerFormRequest struct {
	Year         uint   `json:"year" validate:"required,number"`
	ZoneNo       uint   `json:"zone-no" validate:"required,number"`
	PoleNo       uint   `json:"pole-no" validate:"required,number"`
	ClusterNo    uint   `json:"cluster-no" validate:"required,number"`
	TotalFlowers uint   `json:"total-flowers" validate:"required,number"`
	Condition    string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type FlowerFormResponse struct {
	Message string `json:"message"`
}
