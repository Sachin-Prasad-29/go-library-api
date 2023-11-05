package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachin-prasad-29/go-library-api/controllers"
)

func UserRoutes() *fiber.App {
	app := fiber.New()

	app.Post("/signup", controllers.SignUp)
	app.Post("/usersignin", controllers.UserSignIn)
	app.Post("/adminsignin", controllers.AdminSignIn)
	
	return app
}
