package utils

func CalculateClusterProgress(clusterFormDone, flowerFormDone, pollinationFormDone, podFormDone, preHarvestFormDone bool) uint {
	done := uint(0)

	if clusterFormDone {
		done++
	}
	if flowerFormDone {
		done++
	}
	if pollinationFormDone {
		done++
	}
	if podFormDone {
		done++
	}
	if preHarvestFormDone {
		done++
	}

	return done
}
