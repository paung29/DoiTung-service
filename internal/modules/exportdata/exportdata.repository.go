package exportdata

import "github.com/doitung/DoiTung-service/internal/models"

type ExportDataRepository interface {
	ExportClusterFormsXLSX(yearID uint) ([]models.Cluster, error)
}
