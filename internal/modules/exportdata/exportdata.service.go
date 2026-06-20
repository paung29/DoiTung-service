package exportdata

type ExportDataService interface {
	ExportClusterFormsXLSX(year uint) (ExportXLSXResponse, error)
	ExportHarvestGrading(year int) (ExportXLSXResponse, error)
	ExportHarvestGradingSummary(year int) (ExportXLSXResponse, error)
	ExportStockMovements(year int) (ExportXLSXResponse, error)
}
