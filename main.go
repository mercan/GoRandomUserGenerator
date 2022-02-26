package main

import (
	"log"
	"os"
	"time"

	"github.com/mercan/RandomUserGenerator/router"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mercan/RandomUserGenerator/docs"
)

// @title Go Random User Generator
// @version 1.0.0
// @termsOfService https://swagger.io/terms/
// @contact.name İbrahim Can Mercan
// @contact.email imrcn77@mail.com
// @schemes http https
// @host gorandomusergenerator.herokuapp.com
// @BasePath /api
func main() {
	// Fiber Instance
	app := fiber.New(fiber.Config{
		Prefork:     true,
		ReadTimeout: 60 * time.Second,
	})

	router.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Println(app.Listen(":" + port))
}
