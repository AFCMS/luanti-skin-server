package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	flogger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"luanti-skin-server/auth"
	"luanti-skin-server/database"
	"luanti-skin-server/routes"
	"luanti-skin-server/utils"
)

func main() {
	// Check for Oxipng installation
	if utils.ConfigOptipngEnabled {
		oxipngPresent := utils.OxipngPresent()
		if oxipngPresent {
			log.Println("Oxipng found")
		} else {
			log.Fatalln("Oxipng not found")
		}
	}

	// Connection to Database
	log.Println("Connecting to Database...")
	database.ConnectDB()

	// Create template engine
	engine := html.New("./", ".gohtml")

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName:       "Minetest Skin Server",
		CaseSensitive: false,
		Views:         engine,
	})

	// Initialize Auth
	log.Println("Initializing Auth...")
	auth.Initialize(app)

	// Enable CORS
	app.Use(cors.New())

	// Log requests
	app.Use(flogger.New())

	// Compress responses
	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8080"))
}
