package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	flogger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"luanti-skin-server/backend/auth"
	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/routes"
	"luanti-skin-server/common/oxipng"
)

func main() {
	// Check for Oxipng installation
	oxipngPresent := oxipng.OxipngPresent()
	if oxipngPresent {
		log.Println("Oxipng found")
	} else {
		log.Fatalln("Oxipng not found")
	}

	// Connection to Database
	log.Println("Connecting to Database...")
	database.ConnectDB()

	// Create template engine
	engine := html.New("./", ".gohtml")

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName:       "Luanti Skin Server",
		CaseSensitive: false,
		Views:         engine,
	})

	// Initialize Auth
	log.Println("Initializing Auth...")
	auth.Initialize(app)

	// Enable CORS
	app.Use(cors.New())

	// CSRF Protection
	app.Use(csrf.New())

	// Log requests
	app.Use(flogger.New())

	// Compress responses
	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8081"))
}
