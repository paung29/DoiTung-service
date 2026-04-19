package cluster

import (
	"strconv"

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

func (h *ClusterHandler) GetClustersByZone(context *fiber.Ctx) error {

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	zoneNoStr := context.Query("zoneNo")
	if zoneNoStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneNo is required"))
	}

	zoneNo, err := strconv.Atoi(zoneNoStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneNo"))
	}

	response, err := h.service.GetClustersByZone(year, zoneNo)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *ClusterHandler) GetClusterForm(context *fiber.Ctx) error {
	clusterIdStr := context.Query("clusterId")
	if clusterIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("Cluster Id is required"))
	}
	clusterId, err := strconv.Atoi(clusterIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid ClusterId"))
	}

	response, err := h.service.GetClusterFormByClusterId(clusterId)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
