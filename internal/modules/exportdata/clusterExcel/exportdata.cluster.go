package clusterExcel

import (
	"github.com/doitung/DoiTung-service/internal/models"
	excelutil "github.com/doitung/DoiTung-service/internal/utils"
)

func BuildClusterFormsWorkBook(clusters []models.Cluster) ([]byte, error) {
	sheets := buildClusterSheets(clusters)
	return excelutil.BuildWorkBook(sheets)
}

func buildClusterSheets(clusters []models.Cluster) []excelutil.Sheet {
	groups := excelutil.GroupByZone(clusters, func(cluster models.Cluster) models.Zone {
		return cluster.Pole.Zone
	})

	if len(groups) == 0 {
		return []excelutil.Sheet{
			{
				Name: "No Data",
				Rows: [][]interface{}{clusterHeaderRow()},
			},
		}
	}

	sheets := make([]excelutil.Sheet, 0, len(groups))

	for _, group := range groups {
		rows := [][]interface{}{clusterHeaderRow()}

		for _, cluster := range group.Items {
			rows = append(rows, clusterRow(cluster))
		}

		sheets = append(sheets, excelutil.Sheet{
			Name: group.Zone.ZoneName,
			Rows: rows,
		})
	}

	return sheets
}

func clusterRow(cluster models.Cluster) []interface{} {
	clusterForm := excelutil.FirstOrZero(cluster.ClusterForms)
	flowerForm := excelutil.FirstOrZero(cluster.FlowerForms)
	pollinationForm := excelutil.FirstOrZero(cluster.PollinationForms)
	podForm := excelutil.FirstOrZero(cluster.PodForms)
	preHarvestForm := excelutil.FirstOrZero(cluster.PreHarvestForms)

	return []interface{}{
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
		preHarvestForm.PlantsRemoved,
	}
}

func clusterHeaderRow() []interface{} {
	return []interface{}{
		"Pole Number",
		"Cluster Number",
		"Cluster Condition",
		"Flower Condition",
		"Pollination Condition",
		"Pod Condition",
		"Pre-Harvest Condition",
		"Total Flowers",
		"Unsuccessful Pollination",
		"Pollination Number Pods",
		"Good Flowers",
		"Bad Flowers",
		"Pod Number Pods",
		"Lost Pods",
		"Remaining Pods",
		"Number Pods Round 2",
		"Lost Pods Before Harvest",
		"Removed Pods",
		"Plants Removed",
	}
}
