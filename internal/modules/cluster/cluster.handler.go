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

	zoneIdStr := context.Query("zoneId")
	if zoneIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneId is required"))
	}

	zoneId, err := strconv.Atoi(zoneIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneId"))
	}

	response, err := h.service.GetClustersByZone(year, zoneId)

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

func (h *ClusterHandler) UpdateClusterForm(context *fiber.Ctx) error {
	var form ClusterUpdateRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.UpdateClusterForm(form)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *ClusterHandler) GetClusterFormHistories(context *fiber.Ctx) error {
	userId := context.Locals("account_id").(uint)

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	response, err := h.service.GetClusterFormHistories(userId, uint(year))

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *ClusterHandler) GetAllClustersFormByZone(context *fiber.Ctx) error {
	zoneIdStr := context.Query("zoneId")
	if zoneIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneId is required"))
	}

	zoneId, err := strconv.Atoi(zoneIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneId"))
	}

	if zoneId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("zoneId must be greater than 0"))
	}

	response, err := h.service.GetAllClustersFormByZone(uint(zoneId))

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
func (h *ClusterHandler) GetClusterFilter(context *fiber.Ctx) error {
	zoneIdStr := context.Query("zoneId")
	poleNoStr := context.Query("poleNo")
	clusterNoStr := context.Query("clusterNo")
	progressDoneStr := context.Query("progressDone")

	if zoneIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneId is required"))
	}

	zoneId, err := strconv.Atoi(zoneIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneId"))
	}

	if zoneId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("zoneId must be greater than 0"))
	}

	var poleNo *uint
	if poleNoStr != "" {
		parsedPoleNo, err := strconv.Atoi(poleNoStr)
		if err != nil {
			return utils.HandleError(context, utils.BadRequestError("invalid poleNo"))
		}

		if parsedPoleNo <= 0 {
			return utils.HandleError(context, utils.BadRequestError("poleNo must be greater than 0"))
		}

		poleNoUint := uint(parsedPoleNo)
		poleNo = &poleNoUint
	}

	var clusterNo *uint
	if clusterNoStr != "" {
		parsedClusterNo, err := strconv.Atoi(clusterNoStr)
		if err != nil {
			return utils.HandleError(context, utils.BadRequestError("invalid clusterNo"))
		}

		if parsedClusterNo <= 0 {
			return utils.HandleError(context, utils.BadRequestError("clusterNo must be greater than 0"))
		}

		clusterNoUint := uint(parsedClusterNo)
		clusterNo = &clusterNoUint
	}

	var progressDone *int
	if progressDoneStr != "" {
		parsedProgressDone, err := strconv.Atoi(progressDoneStr)
		if err != nil {
			return utils.HandleError(context, utils.BadRequestError("invalid progressDone"))
		}

		if parsedProgressDone < 0 || parsedProgressDone > 5 {
			return utils.HandleError(context, utils.BadRequestError("progressDone must be between 0 and 5"))
		}

		progressDone = &parsedProgressDone
	}

	response, err := h.service.GetClusterFilter(
		uint(zoneId),
		poleNo,
		clusterNo,
		progressDone,
	)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
