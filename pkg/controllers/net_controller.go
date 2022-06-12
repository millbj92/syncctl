package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/net"

	"github.com/millbj92/synctl/pkg/models/connections"
)

func GetNetIOInfo(c *fiber.Ctx) error {
	info, err := net.IOCounters(true)
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
			"data":  info,
		},
	)
}

func GetConntrackInfo(c *fiber.Ctx) error {
	query := new(connections.ConntrackStatRequest)

	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
	}
	info, err := net.ConntrackStats(query.PerCPU)
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
			"data":  info,
		},
	)
}

func GetConnections(c *fiber.Ctx) error {
	query := new(connections.ConnectionsRequest)

	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
	}
	info, err := net.Connections(query.Kind)
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
			"data":  info,
		},
	)
}

func GetInterfaces(c *fiber.Ctx) error {
	info, err := net.Interfaces()
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
			"data":  info,
		},
	)
}
