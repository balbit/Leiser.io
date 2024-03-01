package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var loadCounter int = 0

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// Increment the load counter
		loadCounter++
		// Pass control to the next middleware/handler
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Railway!",
			"loads":   loadCounter,
		})
	})


	app.Listen(getPort())
}
