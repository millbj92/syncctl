package routes

import (
	"github.com/millbj92/synctl/pkg/controllers"
	"github.com/millbj92/synctl/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)


func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// route.Get("/", middleware.Protected(), controllers.Index)
	// route.Get("/healthz", middleware.Protected(), controllers.Healthz)

	// route.Post("/connections", middleware.Protected(), controllers.CreateConnection)
	// route.Get("/connections", middleware.Protected(), controllers.GetConnections)
	// route.Put("/connections", middleware.Protected(), controllers.UpdateConnection)
	// route.Delete("/connections", middleware.Protected(), controllers.DeleteConnection)

	// route.Put("/users", middleware.Protected(), controllers.UpdateUser)
	// route.Delete("/users", middleware.Protected(), controllers.DeleteUser)
	// route.Get("/users", middleware.Protected(), controllers.GetAllUsers)
	route.Post("/users", middleware.Protected(), controllers.CreateUser)

	// route.Get("/settings", middleware.Protected(), controllers.GetSettings)
	// route.Put("/settings", middleware.Protected(), controllers.UpdateSettings)
	// route.Post("/settings", middleware.Protected(), controllers.CreateSettings)

	// route.Post("/tasks", middleware.Protected(), controllers.CreateTask)
	// route.Get("/tasks", middleware.Protected(), controllers.GetSchedules)
	// route.Put("/tasks", middleware.Protected(), controllers.UpdateTask)
	// route.Delete("/tasks", middleware.Protected(), controllers.DeleteTask)
	// route.Post("/tasks/run", middleware.Protected(), controllers.RunTask)
	// route.Post("/tasks/stop", middleware.Protected(), controllers.StopTask)
	// route.Post("/tasks/pause", middleware.Protected(), controllers.PauseTask)
	// route.Post("/tasks/resume", middleware.Protected(), controllers.ResumeTask)
	// route.Post("/tasks/restart", middleware.Protected(), controllers.RestartTask)
	// route.Get("/tasks/status", middleware.Protected(), controllers.GetTaskStatus)
	// route.Get("/tasks/logs", middleware.Protected(), controllers.GetTaskLogs)
    // route.Get("/tasks/logs/:id/download", middleware.Protected(), controllers.DownloadTaskLogs)

	// route.Post("/links", middleware.Protected(), controllers.CreateLink)
	// route.Get("/links", middleware.Protected(), controllers.GetLinks)
	// route.Put("/links", middleware.Protected(), controllers.UpdateLink)
	// route.Delete("/links", middleware.Protected(), controllers.DeleteLink)

	// route.Post("/roles", middleware.Protected(), controllers.CreateRole)
	// route.Get("/roles", middleware.Protected(), controllers.GetRoles)
	// route.Put("/roles", middleware.Protected(), controllers.UpdateRole)
	// route.Delete("/roles", middleware.Protected(), controllers.DeleteRole)

	// route.Post("/permissions", middleware.Protected(), controllers.CreatePermission)
	// route.Get("/permissions", middleware.Protected(), controllers.GetPermissions)
	// route.Put("/permissions", middleware.Protected(), controllers.UpdatePermission)
	// route.Delete("/permissions", middleware.Protected(), controllers.DeletePermission)

	// route.Post("/remote-configs", middleware.Protected(), controllers.CreateRemoteConfig)
	// route.Get("/remote-configs", middleware.Protected(), controllers.GetRemoteConfigs)
	// route.Put("/remote-configs", middleware.Protected(), controllers.UpdateRemoteConfig)
	// route.Delete("/remote-configs", middleware.Protected(), controllers.DeleteRemoteConfig)
}
