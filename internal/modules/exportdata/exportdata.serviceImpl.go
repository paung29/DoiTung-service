package exportdata

import (
	"errors"
	"fmt"

	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	yearRepo year.YearRepository

	exportDataRepository ExportDataRepository
}

func NewExportDataService(yearRepo year.YearRepository, exportDataRepository ExportDataRepository) ExportDataService {
	return &service{
		yearRepo:             yearRepo,
		exportDataRepository: exportDataRepository,
	}
}

func (s *service) ExportClusterFormsXLSX(year uint) (ExportClusterFormsXLSXResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ExportClusterFormsXLSXResponse{}, utils.NotFoundError("Year not found")
		}
		return ExportClusterFormsXLSXResponse{}, utils.SystemError("Failed to retrieve year information")
	}

	yearID := yearRecord.YearID

	clusterForms, err := s.exportDataRepository.ExportClusterFormsXLSX(yearID)
	if err != nil {
		return ExportClusterFormsXLSXResponse{}, utils.SystemError("Failed to Retrieve export cluster forms")
	}

	if len(clusterForms) == 0 {
		return ExportClusterFormsXLSXResponse{}, utils.NotFoundError("No cluster forms found for the specified year")
	}

	fileBytes, err := BuildFormsWorkBook(clusterForms)
	if err != nil {
		return ExportClusterFormsXLSXResponse{}, utils.SystemError("Failed to generate Excel file")
	}

	return ExportClusterFormsXLSXResponse{
		FileBytes: fileBytes,
		FileName:  fmt.Sprintf("cluster-forms-%d.xlsx", yearRecord.Year),
	}, nil
}
