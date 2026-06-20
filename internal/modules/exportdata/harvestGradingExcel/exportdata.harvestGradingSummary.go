package harvestGradingExcel

import (
	"github.com/doitung/DoiTung-service/internal/models"
	excelutil "github.com/doitung/DoiTung-service/internal/utils"
)

type summaryTotal struct {
	APlusCount       int
	APlusWeight      float64
	ACount           int
	AWeight          float64
	BCount           int
	BWeight          float64
	CCount           int
	CWeight          float64
	DPlusCount       int
	DPlusWeight      float64
	UndersizedCount  int
	UndersizedWeight float64
	RottenCount      int
	RottenWeight     float64
}

func BuildSummaryWorkBook(
	forms []models.HarvestGradingForm,
) ([]byte, error) {
	groups := excelutil.GroupByZone(
		forms,
		func(form models.HarvestGradingForm) models.Zone {
			return form.Pole.Zone
		},
	)

	rows := [][]interface{}{summaryHeaderRow()}

	for _, group := range groups {
		total := calculateTotal(group.Items)

		rows = append(rows, summaryRow(
			group.Zone.ZoneNo,
			group.Zone.ZoneName,
			len(group.Items),
			total,
		))
	}

	if len(groups) > 0 {
		firstDataRow := 2
		lastDataRow := len(rows)

		rows = append(
			rows,
			overallTotalRow(firstDataRow, lastDataRow),
		)
	}

	return excelutil.BuildWorkBook([]excelutil.Sheet{
		{
			Name: "Final Summary",
			Rows: rows,
		},
	})
}

func calculateTotal(
	forms []models.HarvestGradingForm,
) summaryTotal {
	var total summaryTotal

	for _, form := range forms {
		total.APlusCount += form.GradeAPlusCount
		total.APlusWeight += form.GradeAPlusWeight
		total.ACount += form.GradeACount
		total.AWeight += form.GradeAWeight
		total.BCount += form.GradeBCount
		total.BWeight += form.GradeBWeight
		total.CCount += form.GradeCCount
		total.CWeight += form.GradeCWeight
		total.DPlusCount += form.GradeDPlusCount
		total.DPlusWeight += form.GradeDPlusWeight
		total.UndersizedCount += form.UndersizedCount
		total.UndersizedWeight += form.UndersizedWeight
		total.RottenCount += form.RottenCount
		total.RottenWeight += form.RottenWeight
	}

	return total
}

func summaryRow(
	zoneNumber int,
	zoneName string,
	polesRecorded int,
	total summaryTotal,
) []interface{} {
	gradedCount :=
		total.APlusCount +
			total.ACount +
			total.BCount +
			total.CCount +
			total.DPlusCount +
			total.UndersizedCount

	gradedWeight :=
		total.APlusWeight +
			total.AWeight +
			total.BWeight +
			total.CWeight +
			total.DPlusWeight +
			total.UndersizedWeight

	return []interface{}{
		zoneNumber,
		zoneName,
		polesRecorded,
		total.APlusCount,
		total.APlusWeight,
		total.ACount,
		total.AWeight,
		total.BCount,
		total.BWeight,
		total.CCount,
		total.CWeight,
		total.DPlusCount,
		total.DPlusWeight,
		total.UndersizedCount,
		total.UndersizedWeight,
		total.RottenCount,
		total.RottenWeight,
		gradedCount,
		gradedWeight,
		gradedCount + total.RottenCount,
		gradedWeight + total.RottenWeight,
	}
}

func summaryHeaderRow() []interface{} {
	return []interface{}{
		"Zone Number",
		"Zone Name",
		"Poles Recorded",
		"Grade A+ Count",
		"Grade A+ Weight (g)",
		"Grade A Count",
		"Grade A Weight (g)",
		"Grade B Count",
		"Grade B Weight (g)",
		"Grade C Count",
		"Grade C Weight (g)",
		"Grade D+ Count",
		"Grade D+ Weight (g)",
		"Undersized Count",
		"Undersized Weight (g)",
		"Rotten Count",
		"Rotten Weight (g)",
		"Graded Count",
		"Graded Weight (g)",
		"Collected Count",
		"Collected Weight (g)",
	}
}

func overallTotalRow(firstRow, lastRow int) []interface{} {
	row := make([]interface{}, 21)

	row[0] = ""
	row[1] = "OVERALL TOTAL"

	// Zone name is column B, so totals begin at column C.
	for columnNumber := 3; columnNumber <= 21; columnNumber++ {
		columnName, err := excelutil.ColumnName(columnNumber)
		if err != nil {
			continue
		}

		row[columnNumber-1] = excelutil.SumRange(
			columnName,
			firstRow,
			lastRow,
		)
	}

	return row
}
