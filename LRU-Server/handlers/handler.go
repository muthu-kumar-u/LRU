package handlers

import (
	"LRU/data"
	"LRU/lrucontroller"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllCache(cacheController *lrucontroller.LRUCacheController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get all cache entries from the cache controller
		caches := cacheController.GetAllCacheEntries()

		// Transform cache entries into an array of objects
		var data []map[string]interface{}
		for key, value := range caches {
			entry := map[string]interface{}{key: value}
			data = append(data, entry)
		}

		// Return the transformed cache entries as a JSON response
		return c.Status(fiber.StatusOK).JSON(data)
	}
}

func PostCache(cacheController *lrucontroller.LRUCacheController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(data.PostRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		cacheController.SetCacheEntry(req.Key, req.Value, time.Duration(req.ExpireSec)*time.Second)

		// entries := cacheController.GetAllCacheEntries()
		return c.JSON(fiber.Map{
			"message": "Key/Value set in cache",
		})
	}
}

func GetValueFromCache(cacheController *lrucontroller.LRUCacheController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Query("key")

		if key == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "key is missing from the request",
			})
		}

		data, err := cacheController.GetCacheUsingKey(key)

		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(data)
	}
}
