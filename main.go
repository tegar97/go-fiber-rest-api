package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger" // new
	"go-fiber-rest-api/database"
	"go-fiber-rest-api/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.ConnectDatabase()

	routes.SetupRoute(app)
	app.Listen((":8000"))

}
