package customer

type CustomerService interface {
	CreateCustomer(request CreateCustomerRequest) (CreateCustomerResponse, error)
}
