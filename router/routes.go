package router

import (
	"MovieBack/internal/control"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	app.Get("/", control.GetMovies)
	// app.Get("/", control.GetMovies)
	app.Post("/register", control.Register)

}
