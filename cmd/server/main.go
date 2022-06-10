package main

import (
	"os"

	_ "fmt"

	//"github.com/millbj92/synctl/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/millbj92/synctl/internal/configs"
	"github.com/millbj92/synctl/pkg/middleware"
	"github.com/millbj92/synctl/pkg/routes"

	//"github.com/Masterminds/sprig"
	//"html/template"

	//"github.com/gofiber/fiber/v2"
	//"github.com/sirupsen/logrus"
	_ "github.com/joho/godotenv/autoload"
)

// @title Synctl Server
// @version 1.0
// @description The Synctl Server is meant to act as a bridge between the
// @description nodes in a synctl cluster. It will be used to sync files
// @description between nodes, schedule remote nodes for self cleanup, and
// @description perform other general housekeeping tasks.
// @contact.name Brandon Miller
// @contact.email brandon@brandonmiller.io
// @license.name MIT
// @license.url https://github.com/millbj92/synctl/blob/master/LICENSE
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Start() error {
	config := configs.ConfigureFiber()

	app := fiber.New(config)

	middleware.LoadMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)


	if os.Getenv("APP_ENV") == "development" {
		err := app.Listen(":8080")
		if err != nil {
			return err
		}
		} else {
			err := app.Listen(":" + os.Getenv("PORT"))
			if err != nil {
				return err
			}
	}

	return nil
}
