package pod

type PodFormRequest struct {
	ClusterId uint   `json:"clusterId" validate:"required"`
	LostPods  *uint  `json:"lostPods" validate:"required,number"`
	Condition string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type PodFormResponse struct {
	Message string `json:"message"`
}
