package main

import (
	"log"

	"github.com/daffarmd/task-manager/config"
	"github.com/daffarmd/task-manager/internal/handler"
	"github.com/daffarmd/task-manager/internal/middleware"
	"github.com/daffarmd/task-manager/internal/repository"
	"github.com/daffarmd/task-manager/internal/service"
	"github.com/daffarmd/task-manager/pkg/db"
	_ "github.com/daffarmd/task-manager/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load .env
	config.LoadEnv()

	// Connect to DB
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Init repository
	userRepo := repository.NewUserRepository(conn)

	// Init service
	userService := service.NewUserService(userRepo)

	// Init handler
	authHandler := handler.NewAuthHandler(userService)

	// Routes
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Get("/me", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.JSON(fiber.Map{"user_id": userID})
	})

	app.Listen(":8080")
}
