package control

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	fmt.Print("getMovies")
	return c.JSON("Movies List")
}
