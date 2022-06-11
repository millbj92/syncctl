package controllers

import (
	//"context"
	//"time"

	"github.com/millbj92/synctl/app/models/disk"
	//"github.com/millbj92/synctl/pkg/utils"
	//"github.com/millbj92/synctl/platform/cache"
	//"github.com/millbj92/synctl/platform/database"
	"github.com/millbj92/synctl/pkg/monitoring"

	"github.com/gofiber/fiber/v2"
	//"github.com/google/uuid"
)

// GetDiskUsage method: to get a systems disk usage
// @Description Get a systems disk usage
// @Summary Get a systems disk usage
// @Tags disk
// @Accept json
// @Produce json
// @Success 200 {object} DiskResponse
// @Router /api/v1/disk/usage/ [get]
func GetDiskUsage(c *fiber.Ctx) error {

	usage, error := monitoring.DiskUsage("/")
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		&disk.DiskResponse{
			Error: false,
			Msg:   "",
			Data: &disk.DiskUsage{
				Path:        usage.Path,
				Fstype:      usage.Fstype,
				Total:       usage.Total,
				Free:        usage.Free,
				Used:        usage.Used,
				UsedPercent: usage.UsedPercent,
			},
		},
	)
}
