package stockMovementExcel

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	excelutil "github.com/doitung/DoiTung-service/internal/utils"
)

func BuildStockMovementWorkBook(
	movements []models.StockMovement,
) ([]byte, error) {
	rows := [][]interface{}{headerRow()}

	for _, movement := range movements {
		rows = append(rows, movementRow(movement))
	}

	return excelutil.BuildWorkBook([]excelutil.Sheet{
		{
			Name: "Stock Movements",
			Rows: rows,
		},
	})
}

func headerRow() []interface{} {
	return []interface{}{
		"Date",
		"Details",
		"Grade",
		"Production Year",
		"Warehouse",
		"Category",
		"Amount (grams)",
		"Amount (pods)",
		"Customer",
	}
}

func movementRow(movement models.StockMovement) []interface{} {
	return []interface{}{
		movement.RecordedDate.Format("2006-01-02"),
		valueOrBlank(movement.Details),
		string(movement.Grade),
		productionYear(movement),
		warehouseName(movement),
		string(movement.MovementType),
		valueOrBlank(movement.TotalGrams),
		valueOrBlank(movement.TotalPods),
		customerName(movement),
	}
}

func valueOrBlank[T any](value *T) interface{} {
	if value == nil {
		return ""
	}
	return *value
}

func productionYear(movement models.StockMovement) interface{} {
	if movement.ProductionYear == nil {
		return ""
	}

	return movement.ProductionYear.Year
}

func warehouseName(movement models.StockMovement) string {
	warehouse := movement.ToWarehouse

	if movement.MovementType == enums.MovementIssued {
		warehouse = movement.FromWarehouse
	}

	if warehouse == nil {
		return ""
	}

	return warehouse.WarehouseName
}

func customerName(movement models.StockMovement) string {
	if movement.IssuedToCustomer == nil {
		return ""
	}

	return movement.IssuedToCustomer.CustomerName
}
