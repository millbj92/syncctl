package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"

	"github.com/gofiber/websocket/v2"
)

func LoadMiddleware(a *fiber.App) {
	a.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			log.Println("Websocket connection established")
			return c.Next()
		}
		log.Println("Websocket connection failed")
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	a.Use(
		"/api/*",
		csrf.New(),
		helmet.New(),
		cors.New(),
		logger.New(),
	)
}
