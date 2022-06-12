package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/host"
)

func GetHostInfo(c *fiber.Ctx) error {
	hInfo, err := host.Info()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
	}
	return c.JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  hInfo,
		},
	)
}

func GetUserInfo(c *fiber.Ctx) error {
	uInfo, err := host.Users()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
	}
	return c.JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  uInfo,
		},
	)
}
