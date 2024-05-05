package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/name/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Hello, World! " + name)
	})

	app.Listen(":3000")
}
