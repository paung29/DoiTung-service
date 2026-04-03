package cluster

import (
	"github.com/doitung/DoiTung-service/internal/models"
)

type ClusterRepository interface {
	CreatePole(pole *models.Pole) error
	CreateCluster(cluster *models.Cluster) error
	CreateClusterForm(form *models.ClusterForm) error
	FindPoleByZoneAndPoleNo(zoneId uint, poleNo uint) (*models.Pole, error)
	FindClusterByPoleAndClusterNo(poleId uint, clusterNo uint) (*models.Cluster, error)
	FindClusterFormByClusterId(clusterId uint) (*models.ClusterForm, error)
}
