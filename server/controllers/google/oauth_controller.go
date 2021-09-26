package controllersGoogle

import (
	"log"
	"net/http"
)

type OauthController struct{}

func NewOauthController() *OauthController {
	return &OauthController{}
}

func (c *OauthController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.get(w, r)
	default:
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

// redirect to AuthCodeURL
func (c *OauthController) get(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()
	url := config.AuthCodeURL("state") // ! security issue

	http.Redirect(w, r, url, http.StatusFound)
}
