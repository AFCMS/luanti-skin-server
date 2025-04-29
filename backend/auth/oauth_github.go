package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"

	"luanti-skin-server/backend/utils"
)

type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// GitHubUser represents a GitHub user, incomplete
// Try `gh api user` to see a full response example
type GitHubUser struct {
	Id      int    `json:"id"`
	Login   string `json:"login"`
	Message string `json:"message"`
}

func GitHubRedirectURL() string {
	var cdbURL, err = url.Parse(utils.ConfigOAuthRedirectHost)
	if err != nil {
		log.Fatalln(err)
	}
	cdbURL.Path = "/api/account/providers/github/callback"
	return cdbURL.String()
}

func GitHubAuthorize(c fiber.Ctx) error {
	cdbURL, _ := url.Parse("https://github.com/login/oauth/authorize")
	cdbURL.RawQuery = url.Values{
		"scope":        {"public"},
		"client_id":    {utils.ConfigOAuthGitHubClientID},
		"redirect_uri": {GitHubRedirectURL()},
	}.Encode()

	return c.Redirect().To(cdbURL.String())
}

func GitHubExchange(code string) (string, error) {
	c := client.New()
	req := c.R()

	req.SetHeader(fiber.HeaderContentType, fiber.MIMEMultipartForm)
	req.SetHeader(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.SetFormDataWithMap(map[string]string{
		"client_id":     utils.ConfigOAuthGitHubClientID,
		"client_secret": utils.ConfigOAuthGitHubClientSecret,
		"code":          code,
	})

	resp, err := req.Post("https://github.com/login/oauth/access_token")
	if err != nil {
		return "", err
	}

	result := GitHubTokenResponse{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	return result.AccessToken, nil
}

func GitHubFetchUser(token string) (*GitHubUser, error) {
	c := client.New()
	req := c.R()

	req.SetHeaders(map[string]string{
		fiber.HeaderAuthorization: "Bearer " + token,
		fiber.HeaderAccept:        fiber.MIMEApplicationJSON,
	})
	resp, err := req.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}

	d := new(GitHubUser)

	err = resp.JSON(&d)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(d.Message)
	}

	return d, nil
}

func GitHubCallback(c fiber.Ctx) error {
	code := c.Query("code")

	token, err := GitHubExchange(code)
	if err != nil {
		return err
	}

	user, err := GitHubFetchUser(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "GitHub error", "data": err.Error(),
		})
	}

	fmt.Println(user.Id)

	return c.Redirect().To("/")
}
