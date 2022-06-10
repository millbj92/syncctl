package routes

import (
	"github.com/gofiber/fiber/v2"

    localMiddleware "github.com/millbj92/synctl/pkg/middleware"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")

     route.Get("*", localMiddleware.Protected(), swagger.HandlerDefault)
}
