package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func LoadMiddleware(a *fiber.App) {
	a.Use(
		csrf.New(),
		helmet.New(),
		cors.New(),
		logger.New(),
	)
}
