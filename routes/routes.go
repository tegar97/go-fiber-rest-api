package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest-api/Controller/auth"
	"go-fiber-rest-api/Controller/book"
)

func SetupRoute(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/books", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

	app.Post("/api/v1/register", auth.Register)
	app.Post("/api/v1/login", auth.Login)
	app.Get("/api/v1/user", auth.User)
	app.Post("/api/v1/logout", auth.Logout)

}
