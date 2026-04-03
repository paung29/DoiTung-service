package cluster

type ClusterService interface {
	CreateCluster(form ClusterCreateRequest, userId uint) (ClusterCreateResponse, error)
}
