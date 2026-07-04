package utils

import (
	"math"

	"github.com/doitung/DoiTung-service/internal/models"
)

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

func CountTrue(values ...bool) int {
	count := 0
	for _, value := range values {
		if value {
			count++
		}
	}
	return count
}
func CalculateRate(value int64, total int64) float64 {
	if total <= 0 {
		return 0
	}

	return RoundTwoDecimals(float64(value) / float64(total) * 100)
}

func RoundTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
