package controllers

import (
	//"context"
	//"time"

	//"github.com/millbj92/synctl/app/models"
	//"github.com/millbj92/synctl/pkg/utils"
	//"github.com/millbj92/synctl/platform/cache"
	//"github.com/millbj92/synctl/platform/database"
	"github.com/millbj92/synctl/pkg/monitoring"

	"github.com/gofiber/fiber/v2"
	//"github.com/google/uuid"
)


func GetDiskUsage(c *fiber.Ctx) error {

	usage, error := monitoring.DiskUsage("/"); if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg": error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"msg": "",
			"data": usage,
		},
	)
}



