package control

import (
	"MovieBack/config"
	"MovieBack/internal/types"

	"github.com/gofiber/fiber/v2"
)

func MovieDetails(c *fiber.Ctx) error {
	var data types.MovieIdRequest

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON("id is missing")
	}

	movieDetails, err := config.GetRecords("SELECT * FROM movies WHERE movie_id=$1", types.Slice{data.ID})

	if err != nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	} else {
		return c.Status(200).JSON(movieDetails)

	}

}
