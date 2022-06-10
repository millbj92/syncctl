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


func GetAllMemoryStats(c *fiber.Ctx) error {
	usage, error := monitoring.GetAllMemoryStats(); if error != nil {
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

func GetSwapDevices(c *fiber.Ctx) error {
	devices, error := monitoring.GetSwapDevices(); if error != nil {
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
			"data": devices,
		},
	)
}

func GetSwapUsage(c *fiber.Ctx) error {
	usage, error := monitoring.GetSwapUsage(); if error != nil {
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


func GetMemoryUsage(c *fiber.Ctx) error {
	usage, error := monitoring.GetMemoryUsage(); if error != nil {
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
			"data": fiber.Map{
				"total": usage.Total,
				"used": usage.Used,
				"available": usage.Available,
				"used_percent": usage.UsedPercent,
			},
		},
	)
}


func GetFullMemoryUsage(c *fiber.Ctx) error {
	usage, error := monitoring.GetMemoryUsage(); if error != nil {
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


