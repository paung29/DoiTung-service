package customer

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type service struct {
	db   *gorm.DB
	repo CustomerRepository
}

func NewCustomerService(db *gorm.DB, repo CustomerRepository) CustomerService {
	return &service{
		db:   db,
		repo: repo,
	}
}

func (s *service) CreateCustomer(form CreateCustomerRequest) (CreateCustomerResponse, error) {
	customer := &models.Customer{
		CustomerName: form.CustomerName,
		Note:         form.Note,
	}

	if err := s.repo.CreateNewCustomer(customer); err != nil {
		return CreateCustomerResponse{}, err
	}

	return CreateCustomerResponse{
		Message: "Customer created successfully",
	}, nil
}
