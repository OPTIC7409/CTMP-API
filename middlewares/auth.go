package middlewares

import (
	"github.com/OPTIC7409/CTMP-API/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	c.Locals("userID", claims.UserID)
	return c.Next()
}
