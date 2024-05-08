// main.go

package main

import (
	"LRU/lrucontroller"
	"LRU/route/noauth"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	cacheController := lrucontroller.NewLRUCacheController(5)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	noauth.NoAuth(app.Group("/"), cacheController)

	err := app.Listen(":3004")
	if err != nil {
		panic(err)
	}
	fmt.Println("App Started...")
}
