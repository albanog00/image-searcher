package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"giuseppealbano.dev/img-searcher/internal/db"
	"giuseppealbano.dev/img-searcher/internal/unsplash"
)

func SearchHandler(c *fiber.Ctx) error {
	// Retrieves query from url string
	q := c.Query("q")
	if q == "" {
		c.Status(fiber.StatusBadRequest).SendString("query parameter is required")
	}
	page := c.QueryInt("page")
	perPage := c.QueryInt("per_page")

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
	query := fmt.Sprintf("query=%s&page=%d&per_page=%d", q, page, perPage)
	var result *unsplash.SearchPhotoResult
	res, err := redisClient.Client.Get(db.Ctx, query).Result()
	if err != nil {
		log.Printf("Key `%s` not found in Redis, searching...\n", query)

		// Search images in unsplash api
		unsplashClient := unsplash.NewUnsplash(nil)
		log.Printf("Querying unsplash api...")

		searchOptions := &unsplash.SearchOptions{
			Query:   q,
			Page:    page,
			PerPage: perPage,
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
		json, err := json.Marshal(result)
		if err != nil {
			log.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}

		err = redisClient.Client.Set(db.Ctx, query, json, time.Minute*10).Err()
		if err != nil {
			log.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Error",
				"status":  fiber.StatusInternalServerError,
				"result":  nil,
			})
		}
	} else {
		log.Printf("Key `%s` found in Redis\n", query)

		// Unmarshal result from Redis cache
		if err = json.Unmarshal([]byte(res), &result); err != nil {
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
			"count": len(result.Results),
		},
	})
}
