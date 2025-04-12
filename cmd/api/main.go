package main

import (
	"MovieBack/config"
	"MovieBack/internal/container"
	"MovieBack/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := config.IntializeDB()
	if err != nil {
		panic(err)
	}

	config.RunMigrations()
	app := fiber.New()

	config.InitRedis()
	defer config.CloseDB()
	defer config.CloseRedis()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",      // Your frontend URL (adjust if different)
		AllowMethods:     "GET,POST,PUT,DELETE",        // Allowed methods
		AllowHeaders:     "Content-Type,Authorization", // Allowed headers
		AllowCredentials: true,                         // Allow cookies to be sent
	}))
	c := container.NewContainer(db)
	router.Routes(app, c)

	fmt.Print("port is running in 8000")

	app.Listen(":8000")

}
