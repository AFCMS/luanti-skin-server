package main

import (
	"log"

	"minetest-skin-server/database"
	"minetest-skin-server/routes"

	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Reading Env file...")
	_ = godotenv.Load()

	// Connection to Database
	log.Println("Connecting to Database...")

	database.ConnectDB()

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName: "Minetest Skin Server",
	})

	app.Use(flogger.New())

	// API Routes

	api := app.Group("/api")

	api.Get("/info", routes.Info)

	api_account := api.Group("/account")

	api_account.Post("/register", routes.AccountRegister)
	api_account.Post("/login", routes.AccountLogin)
	api_account.Get("/user", routes.AccountUser)
	api_account.Post("/logout", routes.AccountLogout)

	log.Fatalln(app.Listen(":8080"))
}
