package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/helmet/v2"
)

func LoadMiddleware(a *fiber.App) {
	csrfConfig := csrf.Config{
		KeyLookup: 	"header:X-CSRF-Token",
		CookieName: "_csrf",
		CookieSameSite: "Strict",
		Expiration: 24 * time.Hour,
		KeyGenerator: utils.UUID,
	}
	a.Use(
		csrf.New(csrfConfig),
		helmet.New(),
		cors.New(),
		logger.New(),
	)
}
