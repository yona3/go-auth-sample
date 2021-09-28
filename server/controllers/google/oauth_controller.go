package controllersGoogle

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type OauthController struct{}

type OauthResponse struct {
	Ok  bool   `json:"ok"`
	Url string `json:"url"`
}

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
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce) // ! security issue

	data := OauthResponse{true, url}

	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
