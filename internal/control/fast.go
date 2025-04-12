package control

import (
	"MovieBack/internal/container"

	"github.com/gofiber/fiber/v2"
)

// func FastCheck(c *fiber.Ctx) error {
// 	movies, err := config.GetUsers("SELECT movie_id, movie_name, movie_genre, is_active FROM movies", nil)
// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	return c.Status(200).JSON(movies)

// }

// Updated handler in your controller/route package
func GetMoviesHandler(c *container.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get movies from repository
		movies, err := c.MovieRepo.GetMovies(ctx.Context())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(movies)
	}
}
