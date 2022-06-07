package routes

import (
	"github.com/millbj92/synctl/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")


	route.Get("/login", controllers.UserLogin)
	route.Post("/login", controllers.UserLogin)
}
