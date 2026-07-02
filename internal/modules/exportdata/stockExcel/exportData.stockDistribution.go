package stockMovementExcel

import (
	"github.com/doitung/DoiTung-service/internal/models"
	excelutil "github.com/doitung/DoiTung-service/internal/utils"
)

func BuildCustomerDistributionWorkBook(
	movements []models.StockMovement,
) ([]byte, error) {
	rows := [][]interface{}{customerDistributionHeaderRow()}

	for _, movement := range movements {
		rows = append(rows, customerDistributionRow(movement))
	}

	return excelutil.BuildWorkBook([]excelutil.Sheet{
		{
			Name: "Customer Distribution",
			Rows: rows,
		},
	})
}

func customerDistributionHeaderRow() []interface{} {
	return []interface{}{
		"Date",
		"Year",
		"Customer",
		"Grade",
		"Production Year",
		"Warehouse",
		"Price Per Gram",
		"Amount (grams)",
		"Amount (pods)",
		"Total Price",
		"Details",
	}
}

func customerDistributionRow(movement models.StockMovement) []interface{} {
	return []interface{}{
		movement.RecordedDate.Format("2006-01-02"),
		yearValue(movement),
		customerName(movement),
		string(movement.Grade),
		productionYear(movement),
		warehouseName(movement),
		valueOrBlank(movement.PricePerGram),
		valueOrBlank(movement.TotalGrams),
		valueOrBlank(movement.TotalPods),
		totalPrice(movement),
		valueOrBlank(movement.Details),
	}
}

func yearValue(movement models.StockMovement) interface{} {
	if movement.Year.YearID == 0 {
		return ""
	}

	return movement.Year.Year
}

func totalPrice(movement models.StockMovement) interface{} {
	if movement.PricePerGram == nil || movement.TotalGrams == nil {
		return ""
	}

	return float64(*movement.PricePerGram) * *movement.TotalGrams
}
