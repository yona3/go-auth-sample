package controllersGoogle

import "net/http"

type OauthController struct{}

func NewOauthController() *OauthController {
	return &OauthController{}
}

// redirect to AuthCodeURL
func (c *OauthController) Get(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()
	url := config.AuthCodeURL("state") // ! security issue

	http.Redirect(w, r, url, http.StatusFound)
}
