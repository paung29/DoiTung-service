package customer

type CustomerService interface {
	CreateCustomer(form CreateCustomerRequest) (CreateCustomerResponse, error)
	GetAllCustomers() (AllCustomersResponse, error)
	GetCustomerByID(customerID uint) (GetCustomerByIDResponse, error)
	UpdateCustomer(form UpdateCustomerRequest) (UpdateCustomerResponse, error)
}
