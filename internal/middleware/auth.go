package middleware

import (
	"github.com/daffarmd/task-manager/pkg/jwtutil"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if len(auth) < 8 || auth[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		token := auth[7:]
		userID, err := jwtutil.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		c.Locals("user_id", userID)
		return c.Next()
	}
}
