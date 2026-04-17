package cluster

type ClusterService interface {
	CreateCluster(form ClusterCreateRequest, userId uint) (ClusterCreateResponse, error)
	GetClustersByZone(year int, zoneNo int) (ClustersByZoneResponse, error)
}
