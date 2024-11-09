package controllers

import "github.com/gofiber/fiber/v2"

func CreateTask(c *fiber.Ctx) error {
	// TODO: Implement task creation logic
	return c.SendString("Create task endpoint")
}

func GetTasks(c *fiber.Ctx) error {
	// TODO: Implement get tasks logic
	return c.SendString("Get tasks endpoint")
}
