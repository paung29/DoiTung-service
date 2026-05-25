package account

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

func (h *AccountHandler) CreateAccount(context *fiber.Ctx) error {
	var form AccountCreateForm

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateAccount(form)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *AccountHandler) UpdateAccountInfo(context *fiber.Ctx) error {
	var form AccountUpdateInfoForm

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.UpdateAccountInfo(form)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *AccountHandler) UpdateAccountPassword(context *fiber.Ctx) error {
	var form AccountPasswordUpdateForm

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.UpdatePassword(form)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *AccountHandler) GetAllAccounts(context *fiber.Ctx) error {
	response, err := h.service.GetAllAccounts()

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *AccountHandler) GetAccountById(context *fiber.Ctx) error {
	userIdStr := context.Query("userId")
	if userIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("userId is required"))
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid user id"))
	}

	response, err := h.service.GetAccountById(uint(userId))

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *AccountHandler) GetUserAccount(context *fiber.Ctx) error {
	var userId uint = context.Locals("account_id").(uint)
	response, err := h.service.GetAccountById(userId)
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusOK).JSON(response)
}
