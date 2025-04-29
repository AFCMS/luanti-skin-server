package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// ConfigFrontendDevMode is true if the frontend is in development mode (served externally and proxied by the backend)
	ConfigFrontendDevMode bool
	// ConfigFrontendURL is the URL of the frontend when in development mode
	ConfigFrontendURL        string
	ConfigDebugDatabase      bool
	ConfigVerificationGoogle string
	ConfigOAuthRedirectHost  string
	// ConfigOAuthContentDB is the ContentDB OAuth enabled
	ConfigOAuthContentDB             bool
	ConfigOAuthContentDBURL          string
	ConfigOAuthContentDBClientID     string
	ConfigOAuthContentDBClientSecret string
	ConfigOAuthGitHub                bool
	ConfigOAuthGitHubClientID        string
	ConfigOAuthGitHubClientSecret    string
)

func loadConfig() {
	var str string
	var isPresent bool

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_FRONTEND_DEV_MODE")
	if isPresent {
		ConfigFrontendDevMode = str == "true"
	} else {
		ConfigFrontendDevMode = false
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_FRONTEND_URL")
	if isPresent && ConfigFrontendDevMode {
		ConfigFrontendURL = str
	} else {
		ConfigFrontendURL = ""
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_DATABASE_LOGGING")
	if isPresent {
		ConfigDebugDatabase = str == "true"
	} else {
		ConfigDebugDatabase = false
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_VERIFICATION_GOOGLE_SEARCH_CONSOLE")
	if isPresent {
		ConfigVerificationGoogle = str
	} else {
		ConfigVerificationGoogle = ""
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_OAUTH_REDIRECT_HOST")
	if isPresent {
		ConfigOAuthRedirectHost = str
	} else {
		ConfigOAuthRedirectHost = ""
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_OAUTH_CONTENTDB_URL")
	if isPresent {
		ConfigOAuthContentDBURL = str
	} else {
		ConfigOAuthContentDBURL = "https://content.minetest.net"
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_ID")
	if isPresent {
		ConfigOAuthContentDBClientID = str
	} else {
		ConfigOAuthContentDBClientID = ""
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_SECRET")
	if isPresent {
		ConfigOAuthContentDBClientSecret = str
	} else {
		ConfigOAuthContentDBClientSecret = ""
	}

	ConfigOAuthContentDB = ConfigOAuthContentDBClientID != "" || ConfigOAuthContentDBClientSecret != ""

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_ID")
	if isPresent {
		ConfigOAuthGitHubClientID = str
	} else {
		ConfigOAuthGitHubClientID = ""
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_SECRET")
	if isPresent {
		ConfigOAuthGitHubClientSecret = str
	} else {
		ConfigOAuthGitHubClientSecret = ""
	}

	ConfigOAuthGitHub = ConfigOAuthGitHubClientID != "" || ConfigOAuthGitHubClientSecret != ""
}

func init() {
	log.Println("Loading config...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Cannot load .env file: ", err)
	}
	loadConfig()
}
