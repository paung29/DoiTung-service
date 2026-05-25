package customer

type CustomerService interface {
	CreateCustomer(form CreateCustomerRequest) (CreateCustomerResponse, error)
	GetAllCustomers() (AllCustomersResponse, error)
}
