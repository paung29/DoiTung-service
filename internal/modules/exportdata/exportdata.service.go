package exportdata

type ExportDataService interface {
	ExportClusterFormsXLSX(year uint) (ExportClusterFormsXLSXResponse, error)
	ExportHarvestGrading(year int) (ExportXLSXResponse, error)
	ExportHarvestGradingSummary(year int) (ExportXLSXResponse, error)
}
