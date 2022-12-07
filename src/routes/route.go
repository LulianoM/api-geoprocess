package routes

import (
	"github.com/LulianoM/api-geoprocess/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	api.Get("health", controllers.Healthcheck)

	api.Post("commercials", controllers.CommercialsCreateNew)
	api.Get("commercials", controllers.CommercialsGetAll)
}
