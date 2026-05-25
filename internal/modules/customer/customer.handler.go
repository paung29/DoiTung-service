package customer

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(service CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service: service,
	}
}

func (h *CustomerHandler) CreateCustomer(context *fiber.Ctx) error {
	var request CreateCustomerRequest

	if err := utils.ParseAndValidate(context, &request); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateCustomer(request)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *CustomerHandler) GetAllCustomers(context *fiber.Ctx) error {
	response, err := h.service.GetAllCustomers()

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
