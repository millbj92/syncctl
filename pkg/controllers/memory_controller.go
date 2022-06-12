package controllers

import (
	//"context"
	//"time"

	//"github.com/millbj92/synctl/pkg/utils"
	//"github.com/millbj92/synctl/platform/cache"
	//"github.com/millbj92/synctl/platform/database"
	"github.com/millbj92/synctl/pkg/models/memory"
	"github.com/millbj92/synctl/pkg/monitoring"

	"github.com/gofiber/fiber/v2"
	//"github.com/google/uuid"
)

// GetAllMemoryStats method: to get a systems memory stats
// @Description Get a systems memory stats
// @Summary Get a systems memory stats
// @Tags memory
// @Accept json
// @Produce json
// @Success 200 {object} memory.AllMemoryResponse
// @Router /v1/memory/ [get]
func GetAllMemoryStats(c *fiber.Ctx) error {
	usage, err := monitoring.GetAllMemoryStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&memory.AllMemoryResponse{
				Error: true,
				Msg:   error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		&memory.AllMemoryResponse{
			Error: false,
			Msg:   "",
			Data: &memory.MemoryStats{
				VirtualMemory: &memory.VirtualMemory{
					MemoryStat: memory.MemoryStat{
						Total:       usage.VirtualMemory.MemoryStat.Total,
						Used:        usage.VirtualMemory.MemoryStat.Used,
						UsedPercent: usage.VirtualMemory.MemoryStat.UsedPercent,
					},
					Available: usage.VirtualMemory.Available,
				},
				SwapMemory: &memory.SwapMemory{
					MemoryStat: memory.MemoryStat{
						Total:       usage.SwapMemory.MemoryStat.Total,
						Used:        usage.SwapMemory.MemoryStat.Used,
						UsedPercent: usage.SwapMemory.MemoryStat.UsedPercent,
					},
					Free: usage.SwapMemory.Free,
				},
				SwapDevices: usage.SwapDevices,
			},
		},
	)
}

func GetSwapDevices(c *fiber.Ctx) error {
	devices, err := monitoring.GetSwapDevices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  devices,
		},
	)
}

func GetSwapUsage(c *fiber.Ctx) error {
	usage, err := monitoring.GetSwapUsage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  usage,
		},
	)
}

func GetMemoryUsage(c *fiber.Ctx) error {
	usage, err := monitoring.GetMemoryUsage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data": fiber.Map{
				"total":        usage.Total,
				"used":         usage.Used,
				"available":    usage.Available,
				"used_percent": usage.UsedPercent,
			},
		},
	)
}

func GetFullMemoryUsage(c *fiber.Ctx) error {
	usage, err := monitoring.GetMemoryUsage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"msg":   "",
			"data":  usage,
		},
	)
}
