package router

import (
	"MovieBack/internal/container"
	"MovieBack/internal/control"
	"MovieBack/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, c *container.Container) {

	app.Get("/", control.GetMovies)
	app.Get("/movieDetails", control.GetMovies)
	app.Post("/login", middleware.LoginValidator, control.Login)
	app.Post("/register", control.Register)
	app.Get("/tokenVerify", control.TokenValidation)
	app.Get("/move", control.GetMoviesHandler(c))

}
