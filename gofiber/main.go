package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Static("/", "index.html")

	app.Get("/hi", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/api/register", func(c *fiber.Ctx) {
		fmt.Printf("\n/api%s", c.Params("*"))
		// => /api/register
	})

	// GET /john/70
	app.Get("/:name/:age?", func(c *fiber.Ctx) {
		fmt.Printf("\nName: %s, Age: %s", c.Params("name"), c.Params("age"))
		// => Name: john, Age:
	})

	app.Listen(3000)
}
