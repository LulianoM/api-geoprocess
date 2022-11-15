package main

import (
	"github.com/LulianoM/api-geoprocess/src/database"
	"github.com/LulianoM/api-geoprocess/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	database.Connect()
	database.AutoMigrate()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
