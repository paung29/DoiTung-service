package year

import (
	"github.com/doitung/DoiTung-service/internal/models"
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
)


func (repo repository) Create(year *models.Year) error {
	return commonrepo.Create(repo.db, year)
}