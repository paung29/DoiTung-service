package customer

import (
	"strconv"

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

func (h *CustomerHandler) GetCustomerByID(context *fiber.Ctx) error {
	customerIdStr := context.Query("customer_id")
	if customerIdStr == "" {
		return utils.HandleError(context, utils.ValidationError("customerId query parameter is required", nil))
	}

	customerID, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return utils.HandleError(context, utils.ValidationError("Invalid customer ID", nil))
	}

	response, err := h.service.GetCustomerByID(uint(customerID))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *CustomerHandler) UpdateCustomer(context *fiber.Ctx) error {
	var request UpdateCustomerRequest

	if err := utils.ParseAndValidate(context, &request); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.UpdateCustomer(request)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
