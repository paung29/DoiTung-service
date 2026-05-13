package cluster

type ClusterService interface {
	CreateCluster(form ClusterCreateRequest, userId uint) (ClusterCreateResponse, error)
	GetClustersByZone(year int, zoneNo int) (ClustersByZoneResponse, error)
	GetClusterFormByClusterId(clusterId int) (ClusterFormResponse, error)
	UpdateClusterForm(form ClusterUpdateRequest) (ClusterUpdateResponse, error)
	GetClusterFormHistories(userId uint, year uint) (ClusterFormHistoriesResponse, error)
}
