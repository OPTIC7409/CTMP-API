package main

import (
	"log"

	"github.com/OPTIC7409/CTMP-API/database"
	"github.com/OPTIC7409/CTMP-API/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
