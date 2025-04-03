package control

import (
	"MovieBack/config"
	"MovieBack/internal/types"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {

	secretKey := os.Getenv("JWT_PASSWORD")
	user := types.LoginRequest{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON("parsing error")
	}

	userData, dbError := config.GetRecords("SELECT user_id, email, mobile_no, name,password FROM users WHERE mobile_no=$1 AND is_active=$2", types.Slice{user.MobileNo, true})
	if dbError != nil || len(userData) == 0 {
		return c.Status(400).JSON("invalid mobile no")
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(userData[0]["password"].(string)), []byte(user.Password))
	if passwordCheck != nil {
		return c.Status(400).JSON("password is incorrect")
	}
	claims := jwt.MapClaims{
		"userId":   userData[0]["user_id"],
		"email":    userData[0]["email"],
		"mobileNo": userData[0]["mobile_no"],
		"name":     userData[0]["name"],
		"exp":      time.Now().Add(time.Minute * 2).Unix(),
		"iat":      time.Now().Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		fmt.Print(token)
		return c.Status(400).JSON(err.Error())
	}

	refToken, _ := bcrypt.GenerateFromPassword([]byte(userData[0]["user_id"].(string)), bcrypt.DefaultCost)

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    string(refToken),
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HTTPOnly: true,
		SameSite: "Strict",
		Secure:   false,
	})

	go InsertRefreshToken(string(refToken), userData[0]["user_id"].(string))

	return c.Status(200).JSON(fiber.Map{"status": "success", "token": token})

}
