package cluster

type ClusterCreateRequest struct {
	Year      uint   `json:"year" validate:"required,number"`
	ZoneNo    uint   `json:"zoneNo" validate:"required,number"`
	PoleNo    uint   `json:"poleNo" validate:"required,number"`
	ClusterNo uint   `json:"clusterNo" validate:"required,number"`
	Condition string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type ClusterCreateResponse struct {
	Message string `json:"message"`
}

type GetClustersByZoneRequest struct {
	Year   uint `json:"year" validate:"required,number"`
	ZoneNo uint `json:"zoneNo" validate:"required,number"`
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
	RecordedDate string `json:"recordedDate"`
}

type ClusterFormResponse struct {
	ClusterId uint   `json:"clusterId"`
	Location  string `json:"location"`
	PoleNo    int    `json:"poleNo"`
	ClusterNo int    `json:"clusterNo"`
	Condition string `json:"condition"`
}
