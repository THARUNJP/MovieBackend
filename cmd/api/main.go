package main

import (
	"MovieBack/config"
	"MovieBack/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	config.IntializeDB()
	config.InitRedis()
	defer config.CloseDB()
	defer config.CloseRedis()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",      // Your frontend URL (adjust if different)
		AllowMethods:     "GET,POST,PUT,DELETE",        // Allowed methods
		AllowHeaders:     "Content-Type,Authorization", // Allowed headers
		AllowCredentials: true,                         // Allow cookies to be sent
	}))
	router.Routes(app)

	fmt.Print("port is running in 8000")

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal("Error in running server:", err)
	}

}
