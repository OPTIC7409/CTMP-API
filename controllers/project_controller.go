package controllers

import "github.com/gofiber/fiber/v2"

func CreateProject(c *fiber.Ctx) error {
	// TODO: Implement project creation logic
	return c.SendString("Create project endpoint")
}

func GetProjects(c *fiber.Ctx) error {
	// TODO: Implement get projects logic
	return c.SendString("Get projects endpoint")
}
