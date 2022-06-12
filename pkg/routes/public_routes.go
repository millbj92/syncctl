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

	//Memory
	route.Get("/memory/usage", controllers.GetMemoryUsage)
	route.Get("/swap/usage", controllers.GetSwapUsage)
	route.Get("/swap/devices", controllers.GetSwapDevices)
	route.Get("/memory", controllers.GetAllMemoryStats)

	//Cpu
	route.Get("/cpu", controllers.GetCPUInfo)
	route.Get("/cpu/load", controllers.GetCpuLoad)

	//Host
	route.Get("/host", controllers.GetHostInfo)
	route.Get("/host/users", controllers.GetUserInfo)

	//Net
	route.Get("/net/io", controllers.GetNetIOInfo)
	route.Get("/net/iface", controllers.GetInterfaces)
	//route.Get("/net/iface/{iface}", controllers.GetInterfaceById)
	route.Get("/net/conntrack:per_cpu", controllers.GetConntrackInfo)
	route.Get("/net/connections:kind", controllers.GetConnections)
	//Get Connection By ID
}
