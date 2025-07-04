package main

import (
	"os"

	"github.com/daffarmd/task-manager/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Task Manager!")
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
