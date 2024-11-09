package routes

import (
	"github.com/OPTIC7409/CTMP-API/controllers"
	"github.com/OPTIC7409/CTMP-API/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)

	// Protected routes
	api := app.Group("/api", middlewares.AuthMiddleware)
	api.Post("/projects", controllers.CreateProject)
	api.Get("/projects", controllers.GetProjects)
	api.Post("/tasks", controllers.CreateTask)
	api.Get("/tasks", controllers.GetTasks)
}
