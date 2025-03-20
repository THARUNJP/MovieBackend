package main

import (
	"MovieBack/config"
	"MovieBack/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	config.IntializeDB()
	defer config.CloseDB()

	app.Use(cors.New())

	router.Routes(app)

	fmt.Print("port is running in 8000")

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal("Error in running server:", err)
	}

}
