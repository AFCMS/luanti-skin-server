package routes

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/gofiber/fiber/v3/middleware/static"

	"luanti-skin-server/backend/auth"
	"luanti-skin-server/backend/middleware"
	"luanti-skin-server/backend/models"
	"luanti-skin-server/backend/utils"
)

//go:embed all:dist
var distFS embed.FS

func SetupRoutes(app *fiber.App) {
	// API Routes
	api := app.Group("/api")

	api.Get("/info", Info)

	// API Authentication
	apiAccount := api.Group("/account")

	apiAccount.Post("/register", AccountRegister)
	apiAccount.Post("/login", AccountLogin)
	apiAccount.Get("/user", AccountUser, middleware.AuthHandler)
	apiAccount.Post("/logout", AccountLogout, middleware.AuthHandler)

	apiOauthEndpoints := apiAccount.Group("/providers")

	if utils.ConfigOAuthContentDB {
		apiOauthEndpoints.Get("/contentdb", auth.ContentDBAuthorize, middleware.AuthHandlerOptional)
		apiOauthEndpoints.Get("/contentdb/callback", auth.ContentDBCallback, middleware.AuthHandlerOptional)
		apiOauthEndpoints.Post("/contentdb/unlink", auth.ContentDBUnlink, middleware.AuthHandler)
	}

	if utils.ConfigOAuthGitHub {
		apiOauthEndpoints.Get("/github", auth.GitHubAuthorize, middleware.AuthHandlerOptional)
		apiOauthEndpoints.Get("/github/callback", auth.GitHubCallback, middleware.AuthHandlerOptional)
	}

	// Interacting with skins
	apiSkin := api.Group("/skins")

	apiSkin.Get("/list", SkinList)
	apiSkin.Get("/:uuid<guid>", SkinDetails)
	apiSkin.Get("/:uuid<guid>/full", SkinFull)
	apiSkin.Get("/:uuid<guid>/head", SkinHead)
	apiSkin.Post("/:uuid<guid>/delete", NotImplemented, middleware.AuthHandler)
	apiSkin.Post("/create", SkinCreate, middleware.AuthHandler)

	// Interacting with users
	apiUsers := api.Group("/users")

	apiUsers.Get("/list", UsersList)
	apiUsers.Get("/list/banned", NotImplemented, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))
	apiUsers.Get("/:id<int;min(1)>", UsersID)
	apiUsers.Post("/:id<int;min(1)>/ban", NotImplemented, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))
	apiUsers.Post("/:id<int;min(1)>/unban", NotImplemented, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))
	apiUsers.Post("/:id<int;min(1)>/delete", NotImplemented, middleware.AuthHandler)
	apiUsers.Post("/:id<int;min(1)>/permissions", UsersPermissions, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))

	// Handle 404 errors
	api.All("*", NotFound)

	// Serve the React frontend
	if utils.ConfigFrontendDevMode {
		app.Get("*", proxy.Balancer(proxy.Config{
			Servers: []string{utils.ConfigFrontendURL},
			ModifyResponse: func(c fiber.Ctx) error {
				if c.Response().StatusCode() == fiber.StatusNotFound {
					return c.Status(fiber.StatusOK).Render("index", fiber.Map{
						"DevMode": utils.ConfigFrontendDevMode,
					})
				}
				return nil
			},
		}))
	} else {
		distFS, err := fs.Sub(distFS, "dist")
		if err != nil {
			log.Fatalln("Failed to access dist:", err)
		}

		// Parse JSON Vite manifest
		manifest := utils.ViteManifest{}
		data, err := fs.ReadFile(distFS, ".vite/manifest.json")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &manifest)
		if err != nil {
			log.Fatal(err)
		}

		app.Use(static.New("/", static.Config{
			FS:       distFS,
			Compress: true,
		}))

		app.Get("/*", func(c fiber.Ctx) error {
			return c.Render("index", fiber.Map{
				"DevMode":                false,
				"MainCSS":                manifest["src/main.tsx"].Css[0],
				"MainJS":                 manifest["src/main.tsx"].File,
				"GoogleSiteVerification": utils.ConfigVerificationGoogle,
			})
		})
	}
}
