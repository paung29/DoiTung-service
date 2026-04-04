package cluster

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ClusterHandler struct {
	service ClusterService
}

func NewClusterHandler(service ClusterService) *ClusterHandler {
	return &ClusterHandler{service: service}
}

func (h *ClusterHandler) CreateCluster(context *fiber.Ctx) error {
	var userId uint = context.Locals("account_id").(uint)
	var form ClusterCreateRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateCluster(form, userId)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
