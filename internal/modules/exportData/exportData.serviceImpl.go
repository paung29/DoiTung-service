package exportdata

type service struct {
	exportDataRepository ExportDataRepository
}

func NewExportDataService(exportDataRepository ExportDataRepository) ExportDataService {
	return &service{
		exportDataRepository: exportDataRepository,
	}
}

func (s *service) ExportClusterFormsXLSX(year uint) (ExportClusterFormsXLSXResponse, error) {

	return ExportClusterFormsXLSXResponse{}, nil
}
