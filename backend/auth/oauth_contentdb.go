package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"

	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"
	"luanti-skin-server/backend/utils"
)

func CDBRedirectURL() string {
	var cdbURL, err = url.Parse(utils.ConfigOAuthRedirectHost)
	if err != nil {
		log.Fatalln(err)
	}
	cdbURL.Path = "/api/account/providers/contentdb/callback"
	return cdbURL.String()
}

type ContentDBTokenResponse struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Error       string `json:"error"`
}

type ContentDBUser struct {
	Username        string `json:"username"`
	IsAuthenticated bool   `json:"is_authenticated:"`
}

func ContentDBAuthorize(c fiber.Ctx) error {
	cdbURL, _ := url.Parse(utils.ConfigOAuthContentDBURL)
	cdbURL.Path = "/oauth/authorize/"
	cdbURL.RawQuery = url.Values{
		"response_type": {"code"},
		"client_id":     {utils.ConfigOAuthContentDBClientID},
		"redirect_uri":  {CDBRedirectURL()},
	}.Encode()
	log.Println("Authorizing")

	return c.Redirect().To(cdbURL.String())
}

func ContentDBExchange(code string) (string, error) {
	c := client.New()
	req := c.R()

	req.SetHeader(fiber.HeaderContentType, fiber.MIMEMultipartForm)
	req.SetHeader(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.SetFormDataWithMap(map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     utils.ConfigOAuthContentDBClientID,
		"client_secret": utils.ConfigOAuthContentDBClientSecret,
		"code":          code,
	})

	cdbURL, _ := url.Parse(utils.ConfigOAuthContentDBURL)
	cdbURL.Path = "/oauth/token/"

	resp, err := req.Post(cdbURL.String())
	if err != nil {
		return "", err
	}

	result := ContentDBTokenResponse{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	if !result.Success {
		return "", errors.New("CDB: " + result.Error)
	}

	return result.AccessToken, nil
}

func ContentDBFetchUser(token string) (*ContentDBUser, error) {
	c := client.New()
	req := c.R()

	req.SetHeaders(map[string]string{
		fiber.HeaderAuthorization: "Bearer " + token,
		fiber.HeaderAccept:        fiber.MIMEApplicationJSON,
	})
	resp, err := req.Get("https://content.minetest.net/api/whoami/")
	if err != nil {
		return nil, err
	}

	d := new(ContentDBUser)

	err = resp.JSON(&d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func ContentDBCallback(c fiber.Ctx) error {
	code := c.Query("code")

	token, err := ContentDBExchange(code)
	if err != nil {
		return err
	}

	cdbuser, err := ContentDBFetchUser(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "ContentDB error", "data": err.Error(),
		})
	}

	loggedIn := c.Locals("logged_in").(bool)

	if loggedIn {
		user := c.Locals("user").(models.Account)

		if user.CDBUsername != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User already has a ContentDB account",
			})
		} else {
			if err := database.DB.Model(&user).Update("cdb_username", cdbuser.Username).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Database error",
					"data":    err.Error(),
				})
			}
		}
	} else {
		user, err := database.AccountFromCDBUsername(cdbuser.Username)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
				"data":    err.Error(),
			})
		}

		err = InitSession(c, &user)
		if err != nil {
			return err
		}
	}

	return c.Redirect().To("/")
}

func ContentDBUnlink(c fiber.Ctx) error {
	user := c.Locals("user").(models.Account)

	if user.CDBUsername == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User does not have a ContentDB account",
		})
	}

	if err := database.DB.Model(&user).Update("cdb_username", "").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
			"data":    err.Error(),
		})
	}

	return c.Redirect().To("/")
}
