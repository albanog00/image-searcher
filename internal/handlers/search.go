package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

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

	// TODO: Search images

	redisClient, err := db.NewRedisClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"status":  fiber.StatusInternalServerError,
			"result":  nil,
		})
	}

	var result unsplash.PhotosSearch
	res := redisClient.Client.JSONGet(db.Ctx, q, "$").Val()
	if res == "" {
		log.Printf("Key `%s` not found in Redis, searching...\n", q)

		unsplashClient := unsplash.New()
		log.Printf("Querying unsplash api...")

		unsplashRes, err := unsplashClient.Request("GET", fmt.Sprintf("/search/photos/?query=%s", q))
		if err != nil {
			log.Fatalf(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}

		resBody, err := io.ReadAll(unsplashRes.Body)
		if err != nil {
			log.Fatalf(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}

		defer unsplashRes.Body.Close()

		if err := json.Unmarshal(resBody, &result); err != nil {
			log.Fatalf(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}

		// Cache result in redis
		redisClient.Client.JSONSet(db.Ctx, q, "$", result)
	} else {
		log.Printf("Key `%s` found in Redis\n", q)

		var results []unsplash.PhotosSearch
		if err = json.Unmarshal([]byte(res), &results); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}
		result = results[page]
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
