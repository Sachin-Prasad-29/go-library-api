package main

import (
	"github.com/sachin-prasad-29/go-library-api/routes"
)

func main() {
	app := routes.UserRoutes()
	routes.ProductRoutes(app)
	app.Listen(":4000")
}
