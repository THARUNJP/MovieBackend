package control

import (
	"MovieBack/config"
	"MovieBack/internal/types"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

var executeQuery = config.ExecuteQuery

func Register(c *fiber.Ctx) error {
	user := types.UserStruct{}
	err := c.BodyParser(&user)
	fmt.Println(user, "/")
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("err,", user)
	}

	result, Err := executeQuery(`INSERT INTO users(name,email,mobile_no,password) VALUES($1,$2,$3,$4)`,
		types.Slice{user.Name, user.Email, user.MobileNo, user.Password})
	if Err != nil {
		return c.JSON(fiber.Map{"Error": Err})
	}

	fmt.Print(result, "result")

	return c.JSON(fiber.Map{"value": "inserted", "status": "success"})
}
