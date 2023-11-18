package routes

import (
	"github.com/gofiber/fiber/v2"
	"giuseppealbano.dev/img-searcher/internal/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/search", handlers.SearchHandler)
}
