package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/millbj92/synctl/pkg/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/login", controllers.UserLogin)
	route.Post("/login", controllers.UserLogin)

	//change to private routes
	route.Get("/disk/usage", controllers.GetDiskUsage)

	route.Get("/memory/usage", controllers.GetMemoryUsage)
	route.Get("/swap/usage", controllers.GetSwapUsage)
	route.Get("/swap/devices", controllers.GetSwapDevices)
	route.Get("/memory", controllers.GetAllMemoryStats)
}
