package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachin-prasad-29/go-library-api/controllers"
)

func ProductRoutes(app *fiber.App) {
	// Get All Book
	app.Get("/book", controllers.GetAllBooks)

	// Add a Book
	app.Post("/book", controllers.CreateBook)

	// Get A book by Id
	app.Get("/book/:id", controllers.GetBook)

	// Edit a Book by Id
	app.Patch("/book/:id", controllers.EditBook)

	// Delete a Book By Id
	app.Delete("/book/:id", controllers.DeleteBook)

}
