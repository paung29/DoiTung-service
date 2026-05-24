package customer

type CreateCustomerRequest struct {
	CustomerName string  `json:"customer_name" binding:"required"`
	Note         *string `json:"note"`
}

type UpdateCustomerRequest struct {
	CustomerName *string `json:"customer_name"`
	Note         *string `json:"note"`
}

type CreateCustomerResponse struct {
	Message string `json:"message"`
}
