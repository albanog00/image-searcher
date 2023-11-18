package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"giuseppealbano.dev/img-searcher/internal/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
