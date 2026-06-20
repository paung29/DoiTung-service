package harvestGradingExcel

import (
	"github.com/doitung/DoiTung-service/internal/models"
	excelutil "github.com/doitung/DoiTung-service/internal/utils"
)

func BuildHarvestGradingFormsWorkBook(forms []models.HarvestGradingForm) ([]byte, error) {
	sheets := buildHarvestGradingSheets(forms)
	return excelutil.BuildWorkBook(sheets)
}

func buildHarvestGradingSheets(forms []models.HarvestGradingForm) []excelutil.Sheet {
	groups := excelutil.GroupByZone(forms, func(form models.HarvestGradingForm) models.Zone {
		return form.Pole.Zone
	})

	if len(groups) == 0 {
		return []excelutil.Sheet{
			{
				Name: "No Data",
				Rows: [][]interface{}{harvestHeaderRow()},
			},
		}
	}

	sheets := make([]excelutil.Sheet, 0, len(groups))

	for _, group := range groups {
		rows := [][]interface{}{harvestHeaderRow()}

		for _, form := range group.Items {
			rowNumber := len(rows) + 1

			rows = append(rows, harvestRow(form, rowNumber))
		}

		sheets = append(sheets, excelutil.Sheet{
			Name: group.Zone.ZoneName,
			Rows: rows,
		})
	}

	return sheets
}

func harvestHeaderRow(prefix ...string) []interface{} {
	headers := make([]interface{}, 0, len(prefix)+18)
	for _, heading := range prefix {
		headers = append(headers, heading)
	}

	return append(headers,
		"Pole Number",
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
	)
}

func harvestRow(form models.HarvestGradingForm, rowNumber int) []interface{} {

	return []interface{}{
		form.Pole.PoleNo,
		form.GradeAPlusCount,
		form.GradeAPlusWeight,
		form.GradeACount,
		form.GradeAWeight,
		form.GradeBCount,
		form.GradeBWeight,
		form.GradeCCount,
		form.GradeCWeight,
		form.GradeDPlusCount,
		form.GradeDPlusWeight,
		form.UndersizedCount,
		form.UndersizedWeight,
		form.RottenCount,
		form.RottenWeight,

		excelutil.SumCells(rowNumber, "B", "D", "F", "H", "J", "L"),
		excelutil.SumCells(rowNumber, "C", "E", "G", "I", "K", "M"),
		excelutil.SumCells(rowNumber, "P", "N"),
		excelutil.SumCells(rowNumber, "Q", "O"),
	}
}
