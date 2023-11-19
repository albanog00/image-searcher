package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"giuseppealbano.dev/img-searcher/internal/routes"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not retrieve env variables")
	}
}

func main() {
	app := fiber.New()
	// app.Use(recover.New())
	app.Use(logger.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
