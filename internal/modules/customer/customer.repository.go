package customer

import "github.com/doitung/DoiTung-service/internal/models"

type CustomerRepository interface {
	CreateNewCustomer(form *models.Customer) error
	FindAllCustomers() ([]models.Customer, error)
}
