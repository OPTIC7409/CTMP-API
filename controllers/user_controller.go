package controllers

import (
	"github.com/OPTIC7409/CTMP-API/models"
	"github.com/OPTIC7409/CTMP-API/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// TODO: Hash password and save user to database

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func LoginUser(c *fiber.Ctx) error {
	// TODO: Implement login logic
	return c.SendString("Login endpoint")
}
