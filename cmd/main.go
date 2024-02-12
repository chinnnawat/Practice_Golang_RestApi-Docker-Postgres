package main

import (
	"github.com/github.com/chinnnawat/Practice_Golang_RestApi-Docker-Postgres/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
