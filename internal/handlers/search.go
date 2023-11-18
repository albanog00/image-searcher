package handlers

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type SearchResult struct {
	Urls []string `json:"urls"`
}

func SearchHandler(c *fiber.Ctx) error {
	// Retrieves query from url string
	q := c.Query("q")
	if q == "" {
		c.Status(fiber.StatusBadRequest).SendString("query parameter is required")
	}

	// TODO: Search images

	// Simulating images search
	var result SearchResult
	num := rand.Intn(100)
	for i := 0; i < 12; i++ {
		// generate random numbers
		result.Urls = append(
			result.Urls,
			fmt.Sprintf("https://source.unsplash.com/random/?sig=%d", (i%5)*num),
		)
	}

	// json, err := json.Marshal(result)
	// if err != nil {
	// 	log.Println(err)
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Failed to encode result",
	// 		"status":  fiber.StatusBadRequest,
	// 		"result":  nil,
	// 	})
	// }

	// return result as json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Search results",
		"status":  fiber.StatusOK,
		"result": fiber.Map{
			"data":  result,
			"count": len(result.Urls)},
	})
}
