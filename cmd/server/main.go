package main

import (
	_ "fmt"

	//"github.com/millbj92/synctl/pkg/configs"

	//"github.com/Masterminds/sprig"
	//"html/template"

	//"github.com/gofiber/fiber/v2"
	//"github.com/sirupsen/logrus"
	_ "github.com/joho/godotenv/autoload"
)

// func Start() error {
// 	config := configs.ConfigureFiber()

// 	app := fiber.New(config)

// 	middleware.LoadMiddleware(app)

// 	routes.PublicRoutes(app)
// 	routes.PrivateRoutes(app)
// 	routes.NotFoundRoute(app)

// 	if os.Getenv("APP_ENV") == "development" {
// 		err := app.Listen("8080")
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		err := app.Listen(":" + os.Getenv("PORT"))
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
