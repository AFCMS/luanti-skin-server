package auth

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/memory/v2"

	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"
)

var SessionStore *session.Store
var SessionMiddleware fiber.Handler

func Initialize(app *fiber.App) {
	SessionMiddleware, SessionStore = session.NewWithStore(session.Config{
		Storage:           memory.New(), // TODO: allow to use Redis
		KeyLookup:         "cookie:session_id",
		CookieDomain:      "",
		CookiePath:        "",
		CookieSecure:      false, // TODO: handle production
		CookieHTTPOnly:    true,
		CookieSessionOnly: false,
	})

	app.Use(SessionMiddleware)
}

func InitSession(c fiber.Ctx, user *models.Account) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error interacting with session", "data": err.Error()})
	}

	if sess.Fresh() {
		// Get session ID
		sid := sess.ID()

		// Get user ID
		uid := user.ID

		// Save session data
		sess.Set("uid", uid)
		sess.Set("sid", sid)
		sess.Set("ip", c.IP())
		sess.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
		sess.Set("ua", string(c.Request().Header.UserAgent()))

		err := sess.Save()
		if err != nil {
			log.Println(err)
		}
	}

	if err := database.AccountSetLastConnection(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error interacting with database", "data": err.Error()})
	}

	return nil
}
