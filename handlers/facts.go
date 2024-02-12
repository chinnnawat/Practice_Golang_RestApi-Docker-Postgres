package handlers

import (
	"github.com/github.com/chinnnawat/Practice_Golang_RestApi-Docker-Postgres/database"
	"github.com/github.com/chinnnawat/Practice_Golang_RestApi-Docker-Postgres/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Chinnawat wongket :)!")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
