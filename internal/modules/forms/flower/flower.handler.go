package flower

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type FlowerHandler struct {
	service FlowerService
}

func NewFlowerHandler(service FlowerService) *FlowerHandler {
	return &FlowerHandler{service: service}
}

func (h *FlowerHandler) CreateOrUpdateFlowerForm(context *fiber.Ctx) error {

	var userId uint = context.Locals("account_id").(uint)
	var form FlowerFormRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateOrUpdateFlowerForm(form, userId)

	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *FlowerHandler) GetFlowerFormDetails(context *fiber.Ctx) error {

	clusterIdStr := context.Query("clusterId")
	if clusterIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("clusterId is required"))
	}

	clusterId, err := strconv.Atoi(clusterIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid clusterId"))
	}

	if clusterId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("clusterId must be a positive integer"))
	}

	response, err := h.service.GetFlowerFormDetailsByClusterID(uint(clusterId))

	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}
