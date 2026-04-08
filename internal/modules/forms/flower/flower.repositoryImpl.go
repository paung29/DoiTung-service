package flower

import (
	commonRepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewFlowerRepository(db *gorm.DB) FlowerRepository {
	return &repository{db: db}
}

// CreateFlowerForm implements [FlowerRepository].
func (r *repository) CreateFlowerForm(db *gorm.DB, form *models.FlowerForm) error {
	return commonRepo.Create(db, form)
}
