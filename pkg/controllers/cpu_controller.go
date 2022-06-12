package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

func GetCPUInfo(c *fiber.Ctx) error {
	v, err := cpu.Info()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": true, "msg": err.Error()})
	}
	return c.JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  v,
		},
	)
}

func GetCpuLoad(C *fiber.Ctx) error {
	v, err := load.Avg()
	if err != nil {
		return C.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
	}
	return C.JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  v,
		},
	)
}
