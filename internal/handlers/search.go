package handlers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"giuseppealbano.dev/img-searcher/internal/db"
	"giuseppealbano.dev/img-searcher/internal/unsplash"
)

func SearchHandler(c *fiber.Ctx) error {
	// Retrieves query from url string
	page := c.QueryInt("page")
	q := c.Query("q")
	if q == "" {
		c.Status(fiber.StatusBadRequest).SendString("query parameter is required")
	}

	q = strings.ToLower(q)

	// Redis Client
	redisClient, err := db.NewRedisClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"status":  fiber.StatusInternalServerError,
			"result":  nil,
		})
	}

	// Search query in Redis cache
	var result *unsplash.SearchPhotoResult
	res := redisClient.Client.JSONGet(db.Ctx, q, "$").Val()
	if res == "" {
		log.Printf("Key `%s` not found in Redis, searching...\n", q)

		// Search images in unsplash api
		unsplashClient := unsplash.NewUnsplash(nil)
		log.Printf("Querying unsplash api...")

		searchOptions := &unsplash.SearchOptions{
			Page:    page,
			PerPage: 0,
			Query:   q,
		}

		result, _, err = unsplashClient.Search.Photos(searchOptions)
		if err != nil {
			log.Fatalf(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}

		// Set result in Redis cache
		redisClient.Client.JSONSet(db.Ctx, q, "$", result)
	} else {
		log.Printf("Key `%s` found in Redis\n", q)

		// Unmarshal result from Redis cache
		var results []*unsplash.SearchPhotoResult
		if err = json.Unmarshal([]byte(res), &results); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}
		result = results[0]
	}

	// return result as json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Search results",
		"status":  fiber.StatusOK,
		"result": fiber.Map{
			"data":  result,
			"count": len(result.Results)},
	})
}
