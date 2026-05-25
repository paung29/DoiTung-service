package customer

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/utils"
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
		return CreateCustomerResponse{}, utils.BadRequestError("Failed to create customer")
	}

	return CreateCustomerResponse{
		Message: "Customer created successfully",
	}, nil
}

func (s *service) GetAllCustomers() (AllCustomersResponse, error) {
	customers, err := s.repo.FindAllCustomers()
	if err != nil {
		return AllCustomersResponse{}, err
	}

	customerDetails := make([]CustomerDetails, len(customers))
	for i, customer := range customers {
		customerDetails[i] = CustomerDetails{
			ID:           int(customer.CustomerID),
			CustomerName: customer.CustomerName,
			Note:         customer.Note,
		}
	}

	return AllCustomersResponse{
		Customers: customerDetails,
	}, nil
}

func (s *service) GetCustomerByID(customerID uint) (GetCustomerByIDResponse, error) {
	customer, err := s.repo.FindByCustomerID(customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return GetCustomerByIDResponse{}, utils.NotFoundError("Customer not found")
		}
		return GetCustomerByIDResponse{}, utils.BadRequestError("Failed to get customer")
	}

	customerDetails := CustomerDetails{
		ID:           int(customer.CustomerID),
		CustomerName: customer.CustomerName,
		Note:         customer.Note,
	}

	return GetCustomerByIDResponse{
		Customer: customerDetails,
	}, nil

}
