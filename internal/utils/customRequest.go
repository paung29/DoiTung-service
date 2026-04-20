package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetClusterIDFromQuery(c *fiber.Ctx) (uint, error) {

	clusterIdStr := c.Query("clusterId")
	if clusterIdStr == "" {
		return 0, BadRequestError("clusterId is required")
	}
	clusterId, err := strconv.Atoi(clusterIdStr)
	if err != nil {
		return 0, BadRequestError("invalid clusterId")
	}
	if clusterId <= 0 {
		return 0, BadRequestError("clusterId must be a positive integer")
	}
	return uint(clusterId), nil

}
