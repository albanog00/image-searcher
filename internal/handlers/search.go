package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"giuseppealbano.dev/img-searcher/internal/db"
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

	redisClient, err := db.NewRedisClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"status":  fiber.StatusInternalServerError,
			"result":  nil,
		})
	}

	var result SearchResult
	res, err := redisClient.Client.Get(db.Ctx, q).Result()
	if err == redis.Nil {
		log.Printf("Key %s not found in Redis, searching...\n", q)

		// Simulating images search
		if len(result.Urls) == 0 {
			num := rand.Intn(100)
			for i := 0; i < 12; i++ {
				// generate random urls
				result.Urls = append(
					result.Urls,
					fmt.Sprintf("https://source.unsplash.com/random/?sig=%d", (i%5)*num),
				)
			}
		}

		// Cache result in redis
		json, err := json.Marshal(result)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}
		redisClient.Client.Set(db.Ctx, q, json, time.Hour*24)
	} else {
		log.Printf("Key %s found in Redis\n", q)

		err = json.Unmarshal([]byte(res), &result)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}
	}

	// return result as json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Search results",
		"status":  fiber.StatusOK,
		"result": fiber.Map{
			"data":  result,
			"count": len(result.Urls)},
	})
}
