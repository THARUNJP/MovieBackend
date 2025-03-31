package control

import (
	"MovieBack/config"
	"MovieBack/internal/types"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := types.UserStruct{}
	err := c.BodyParser(&user)
	fmt.Println(user, "/")
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("err,", user)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON("ERROR HASHING PASSWORD")
	}
	hashedPassword := string(hashed)
	fmt.Println(hashedPassword, "hashed", hashed)

	result, Err := config.ExecuteQuery(`INSERT INTO users(name,email,mobile_no,password) VALUES($1,$2,$3,$4)`,
		types.Slice{user.Name, user.Email, user.MobileNo, hashedPassword})
	if Err != nil {
		return c.JSON(fiber.Map{"Error": Err})
	}

	fmt.Print(result, "result")

	return c.JSON(fiber.Map{"value": "inserted", "status": "success"})
}
