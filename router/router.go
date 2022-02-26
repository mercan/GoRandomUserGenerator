package router

import (
	"github.com/mercan/RandomUserGenerator/handler"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${magenta}${time}${reset}] - Status:${status} Latency:${latency} Method:${method}" +
			"Path:${path} User-Agent:${ua}\n",
	}))

	app.Get("/dashboard", monitor.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/user", handler.GenerateUser)
	v1.Get("/address", handler.GenerateAddress)
	v1.Get("/credit_card", handler.GenerateCreditCard)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
