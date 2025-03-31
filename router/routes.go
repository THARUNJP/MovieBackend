package router

import (
	"MovieBack/internal/control"
	"MovieBack/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	app.Get("/", control.GetMovies)
	app.Post("/movieDetails", control.MovieDetails)
	app.Post("/login", middleware.LoginValidator, control.Login)
	app.Post("/register", control.Register)

}
