package control

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Upcoming(c *fiber.Ctx) {

	fmt.Print("hi from movies")

	c.JSON("welcome to app")

}
