package utils

import "github.com/doitung/DoiTung-service/internal/models"

func CalculateClusterProgress(clusterInfo models.Cluster) uint {
	done := uint(0)

	if clusterInfo.ClusterFormDone {
		done++
	}
	if clusterInfo.FlowerFormDone {
		done++
	}
	if clusterInfo.PollinationFormDone {
		done++
	}
	if clusterInfo.PodFormDone {
		done++
	}
	if clusterInfo.PreHarvestFormDone {
		done++
	}

	return done
}
