package exportdata

type ExportDataService interface {
	ExportClusterFormsXLSX(year uint) (ExportClusterFormsXLSXResponse, error)
}
