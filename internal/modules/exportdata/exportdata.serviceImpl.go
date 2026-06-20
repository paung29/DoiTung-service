package exportdata

import (
	"errors"
	"fmt"
	"time"

	clusterExcel "github.com/doitung/DoiTung-service/internal/modules/exportdata/clusterExcel"
	"github.com/doitung/DoiTung-service/internal/modules/exportdata/harvestGradingExcel"
	stockMovementExcel "github.com/doitung/DoiTung-service/internal/modules/exportdata/stockExcel"
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

func (s *service) ExportClusterFormsXLSX(year uint) (ExportXLSXResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ExportXLSXResponse{}, utils.NotFoundError("Year not found")
		}
		return ExportXLSXResponse{}, utils.SystemError("Failed to retrieve year information")
	}

	yearID := yearRecord.YearID

	clusterForms, err := s.exportDataRepository.ExportClusterFormsXLSX(yearID)
	if err != nil {
		return ExportXLSXResponse{}, utils.SystemError("Failed to Retrieve export cluster forms")
	}

	if len(clusterForms) == 0 {
		return ExportXLSXResponse{}, utils.NotFoundError("No cluster forms found for the specified year")
	}

	fileBytes, err := clusterExcel.BuildClusterFormsWorkBook(clusterForms)
	if err != nil {
		return ExportXLSXResponse{}, utils.SystemError("Failed to generate Excel file")
	}

	return ExportXLSXResponse{
		FileBytes: fileBytes,
		FileName:  fmt.Sprintf("cluster-forms-%d.xlsx", yearRecord.Year),
	}, nil
}

func (s *service) ExportHarvestGrading(
	yearValue int,
) (ExportXLSXResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(yearValue)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ExportXLSXResponse{},
				utils.NotFoundError("year not found")
		}

		return ExportXLSXResponse{},
			utils.SystemError("failed to retrieve year")
	}

	forms, err := s.exportDataRepository.
		FindHarvestGradingFormsByYearID(yearRecord.YearID)
	if err != nil {
		return ExportXLSXResponse{},
			utils.SystemError(
				"failed to retrieve harvest and grading forms",
			)
	}

	if len(forms) == 0 {
		return ExportXLSXResponse{},
			utils.NotFoundError(
				"no harvest and grading forms found for this year",
			)
	}

	fileBytes, err :=
		harvestGradingExcel.BuildHarvestGradingFormsWorkBook(forms)
	if err != nil {
		return ExportXLSXResponse{},
			utils.SystemError("failed to generate Excel file")
	}

	return ExportXLSXResponse{
		FileBytes: fileBytes,
		FileName: fmt.Sprintf(
			"harvest-grading-%d-%s.xlsx",
			yearRecord.Year,
			time.Now().Format("2006-01-02"),
		),
	}, nil
}

func (s *service) ExportHarvestGradingSummary(year int) (ExportXLSXResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(year)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ExportXLSXResponse{},
				utils.NotFoundError("year not found")
		}

		return ExportXLSXResponse{},
			utils.SystemError("failed to retrieve year")
	}

	yearID := yearRecord.YearID

	forms, err := s.exportDataRepository.
		FindHarvestGradingFormsByYearID(yearID)
	if err != nil {
		return ExportXLSXResponse{},
			utils.SystemError(
				"failed to retrieve harvest and grading forms",
			)
	}

	if len(forms) == 0 {
		return ExportXLSXResponse{},
			utils.NotFoundError(
				"no harvest and grading forms found for this year",
			)
	}

	fileBytes, err := harvestGradingExcel.BuildSummaryWorkBook(forms)
	if err != nil {
		return ExportXLSXResponse{},
			utils.SystemError("failed to generate Excel file")
	}
	return ExportXLSXResponse{
		FileBytes: fileBytes,
		FileName: fmt.Sprintf(
			"harvest-grading-summary-%d-%s.xlsx",
			yearRecord.Year,
			time.Now().Format("2006-01-02"),
		),
	}, nil
}

func (s *service) ExportStockMovements(yearValue *int) (ExportXLSXResponse, error) {
	var yearID *uint
	fileName := "stock-movements-all-years"

	if yearValue != nil {
		yearRecord, err := s.yearRepo.FindByYear(*yearValue)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ExportXLSXResponse{},
					utils.NotFoundError("year not found")
			}

			return ExportXLSXResponse{},
				utils.SystemError("failed to retrieve year")
		}

		yearID = &yearRecord.YearID
		fileName = fmt.Sprintf(
			"stock-movements-%d",
			yearRecord.Year,
		)
	}

	movements, err :=
		s.exportDataRepository.FindStockMovements(yearID)
	if err != nil {
		return ExportXLSXResponse{},
			utils.SystemError("failed to retrieve stock movements")
	}

	if len(movements) == 0 {
		return ExportXLSXResponse{},
			utils.NotFoundError("no stock movements found")
	}

	fileBytes, err :=
		stockMovementExcel.BuildStockMovementWorkBook(movements)
	if err != nil {
		return ExportXLSXResponse{},
			utils.SystemError("failed to generate Excel file")
	}

	return ExportXLSXResponse{
		FileBytes: fileBytes,
		FileName: fmt.Sprintf(
			"%s-%s.xlsx",
			fileName,
			time.Now().Format("2006-01-02"),
		),
	}, nil
}
