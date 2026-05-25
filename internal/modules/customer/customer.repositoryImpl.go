package customer

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"

	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &repository{db: db}
}

func (r *repository) CreateNewCustomer(form *models.Customer) error {
	return commonrepo.Create(r.db, form)
}

func (r *repository) FindAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
