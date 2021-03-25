package book

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest-api/database"
	"go-fiber-rest-api/models"
	"net/http"
)

func GetBooks(c *fiber.Ctx) error  {
	var books []models.Book
	database.DB.Find(&books);


	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"succes" : true,
		"data" : books,
	})

}

func GetBook(c *fiber.Ctx) error {
	msg := "Get Book"
	return c.SendString(msg)
}

func NewBook(c *fiber.Ctx) error {
	var input models.CreateBook

	c.BodyParser(&input)

	if len(input.Title) == 0 && len(input.Author) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"status" : false,
			"error" : "Required Title , Author,Rating",
		})
	}


	book  := models.Book{Title: input.Title,Author: input.Author, Rating: input.Rating}
	database.DB.Create(&book)


	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success" : true,
		"data" : book,
	})

}

func DeleteBook(c *fiber.Ctx) error {
	msg := "Delete Books"
	return c.SendString(msg)
}






