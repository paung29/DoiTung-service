package exportdata

import (
	"sort"

	"github.com/doitung/DoiTung-service/internal/models"

	excelutil "github.com/doitung/DoiTung-service/internal/utils"
)

type exportSheet struct {
	Name string
	Rows [][]interface{}
}

func BuildFormsWorkBook(clusters []models.Cluster) ([]byte, error) {
	return excelutil.BuildWorkBook(buildSheets(clusters))
}

func buildSheets(clusters []models.Cluster) []excelutil.Sheet {
	zoneClusters := make(map[uint][]models.Cluster)
	zoneOrder := make([]uint, 0)
	seenZones := make(map[uint]bool)

	for _, cluster := range clusters {
		zoneID := cluster.Pole.Zone.ZoneID
		zoneClusters[zoneID] = append(zoneClusters[zoneID], cluster)

		if !seenZones[zoneID] {
			seenZones[zoneID] = true
			zoneOrder = append(zoneOrder, zoneID)
		}

		sort.Slice(zoneOrder, func(i, j int) bool {
			left := zoneClusters[zoneOrder[i]][0].Pole.Zone.ZoneNo
			right := zoneClusters[zoneOrder[j]][0].Pole.Zone.ZoneNo
			return left < right
		})
	}

	sheets := make([]excelutil.Sheet, 0, len(zoneOrder))

	for _, zoneID := range zoneOrder {
		zone := zoneClusters[zoneID][0].Pole.Zone
		rows := [][]interface{}{exportHeaderRow()}

		for _, cluster := range zoneClusters[zoneID] {
			clusterForm := FirstOrZero(cluster.ClusterForms)
			flowerForm := FirstOrZero(cluster.FlowerForms)
			pollinationForm := FirstOrZero(cluster.PollinationForms)
			podForm := FirstOrZero(cluster.PodForms)
			preHarvestForm := FirstOrZero(cluster.PreHarvestForms)

			rows = append(rows, []interface{}{
				cluster.Pole.PoleNo,
				cluster.ClusterNo,
				string(clusterForm.Condition),
				string(flowerForm.Condition),
				string(pollinationForm.Condition),
				string(podForm.Condition),
				string(preHarvestForm.Condition),
				flowerForm.TotalFlowers,
				pollinationForm.UnsuccessfulPollination,
				pollinationForm.NumberPods,
				pollinationForm.GoodFlowers,
				pollinationForm.BadFlowers,
				podForm.NumberPods,
				podForm.LostPods,
				podForm.RemainingPods,
				preHarvestForm.NumberPodsSecondRound,
				preHarvestForm.LostPodsBeforeHarvest,
				preHarvestForm.RemovedPods,
				preHarvestForm.PlantsRemoved})
		}
		sheets = append(sheets, excelutil.Sheet{Name: zone.ZoneName, Rows: rows})
	}

	if len(sheets) == 0 {
		sheets = append(sheets, excelutil.Sheet{Name: "No Data", Rows: [][]interface{}{exportHeaderRow()}})
	}

	return sheets
}

func exportHeaderRow() []interface{} {
	return []interface{}{
		"Pole Number", "Cluster Number", "Cluster Condition",
		"Flower Condition", "Pollination Condition", "Pod Condition",
		"Pre-Harvest Condition", "Total Flowers", "Unsuccessful Pollination",
		"Pollination Number Pods", "Good Flowers", "Bad Flowers",
		"Pod Number Pods", "Lost Pods", "Remaining Pods",
		"Number Pods Round 2", "Lost Pods Before Harvest",
		"Removed Pods", "Plants Removed",
	}
}

func FirstOrZero[T any](items []T) T {
	var zero T
	if len(items) == 0 {
		return zero
	}
	return items[0]
}
