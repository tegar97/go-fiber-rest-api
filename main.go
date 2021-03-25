package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest-api/book"
	"go-fiber-rest-api/database"
	"github.com/gofiber/fiber/v2/middleware/logger" // new

)


func setupRoute(app *fiber.App){
	app.Get("/api/v1/books",book.GetBooks);
	app.Get("/api/v1/book/:id",book.GetBook);
	app.Post("/api/v1/books",book.NewBook);
	app.Delete("/api/v1/book/:id",book.DeleteBook);

}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database.ConnectDatabase()
	setupRoute(app);



	app.Listen((":8000"))



}

