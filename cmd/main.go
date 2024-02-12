package main

import (
	"github.com/github.com/chinnnawat/Practice_Golang_RestApi-Docker-Postgres/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Chinnawat!")
	})

	app.Listen(":3000")
}
