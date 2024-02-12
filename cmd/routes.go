package main

import (
	"github.com/github.com/chinnnawat/Practice_Golang_RestApi-Docker-Postgres/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
