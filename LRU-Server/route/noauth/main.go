// noauth/route.go

package noauth

import (
	"LRU/handlers"
	"LRU/lrucontroller"

	"github.com/gofiber/fiber/v2"
)

func NoAuth(app fiber.Router, cacheController *lrucontroller.LRUCacheController) {
	app.Get("api/GetKey", handlers.GetAllCache(cacheController))
	app.Get("api/GetValue", handlers.GetValueFromCache(cacheController))
	app.Put("api/PostKey", handlers.PostCache(cacheController))
}
