package cluster

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewClusterRepository(db *gorm.DB) ClusterRepository {
	return &repository{db: db}
}

// CreatePole implements [ClusterRepository].
func (r *repository) CreatePole(pole *models.Pole) error {
	return commonrepo.Create(r.db, pole)
}

// CreateCluster implements [ClusterRepository].
func (r *repository) CreateCluster(cluster *models.Cluster) error {
	return commonrepo.Create(r.db, cluster)
}

// CreateClusterForm implements [ClusterRepository].
func (r *repository) CreateClusterForm(form *models.ClusterForm) error {
	return commonrepo.Create(r.db, form)
}

// FindPoleByZoneAndPoleNo implements [ClusterRepository].
func (r *repository) FindPoleByZoneAndPoleNo(zoneId uint, poleNo uint) (*models.Pole, error) {
	var pole models.Pole
	if err := r.db.Where("zone_id = ? AND pole_no = ?", zoneId, poleNo).First(&pole).Error; err != nil {
		return nil, err
	}
	return &pole, nil
}

// FindClusterByPoleAndClusterNo implements [ClusterRepository].
func (r *repository) FindClusterByPoleAndClusterNo(poleId uint, clusterNo uint) (*models.Cluster, error) {
	var cluster models.Cluster
	if err := r.db.Where("pole_id = ? AND cluster_no = ?", poleId, clusterNo).First(&cluster).Error; err != nil {
		return nil, err
	}
	return &cluster, nil
}

// FindClusterFormById implements [ClusterRepository].
func (r *repository) FindClusterFormByClusterId(clusterId uint) (*models.ClusterForm, error) {
	var form models.ClusterForm
	if err := r.db.Where("cluster_id = ?", clusterId).First(&form).Error; err != nil {
		return nil, err
	}
	return &form, nil
}

func (r *repository) UpdateCluster(cluster *models.Cluster) error {
	return commonrepo.Save(r.db, cluster)
}
