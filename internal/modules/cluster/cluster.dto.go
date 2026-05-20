package cluster

type ClusterCreateRequest struct {
	Year      uint   `json:"year" validate:"required,number"`
	ZoneId    uint   `json:"zoneId" validate:"required,number"`
	PoleNo    uint   `json:"poleNo" validate:"required,number"`
	ClusterNo uint   `json:"clusterNo" validate:"required,number"`
	Condition string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type ClusterCreateResponse struct {
	Message string `json:"message"`
}

type GetClustersByZoneRequest struct {
	Year   uint `json:"year" validate:"required,number"`
	ZoneId uint `json:"zoneId" validate:"required,number"`
}

type ClustersByZoneResponse struct {
	Clusters []ClusterInfo `json:"clusters"`
}

type ClusterInfo struct {
	No           int    `json:"no"`
	ClusterId    uint   `json:"clusterId"`
	Location     string `json:"location"`
	PoleNo       int    `json:"poleNo"`
	ClusterNo    int    `json:"clusterNo"`
	ProgressDone int    `json:"progressDone"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type ClusterFormResponse struct {
	ClusterId uint   `json:"clusterId"`
	Location  string `json:"location"`
	PoleNo    int    `json:"poleNo"`
	ClusterNo int    `json:"clusterNo"`
	Condition string `json:"condition"`
}

type ClusterUpdateRequest struct {
	ClusterId uint   `json:"clusterId" validate:"required,number"`
	Condition string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type ClusterUpdateResponse struct {
	Message string `json:"message"`
}

type ClusterFormHistoriesResponse struct {
	ClusterFormHistories []ClusterInfo `json:"clusterFormHistories"`
}
