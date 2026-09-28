package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/*", static.New("./static"))

	log.Fatal(app.Listen(":" + port))
}
