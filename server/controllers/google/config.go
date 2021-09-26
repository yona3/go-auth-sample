package controllersGoogle

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetConfig() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"openid", "email", "profile"},
		RedirectURL:  "http://localhost:8080/google/callback",
	}

	return config
}
