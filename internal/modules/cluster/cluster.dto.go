package cluster

type ClusterCreateRequest struct {
	YearId    uint   `json:"year-id" validate:"required,number"`
	ZoneNo    uint   `json:"zone-no" validate:"required,number"`
	PoleNo    uint   `json:"pole-no" validate:"required,number"`
	ClusterNo uint   `json:"cluster-no" validate:"required,number"`
	Condition string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type ClusterCreateResponse struct {
	Message string `json:"message"`
}
