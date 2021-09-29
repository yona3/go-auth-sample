package controllersGoogle

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/utils"
	"golang.org/x/oauth2"
)

type OauthController struct {
	state string
}

type GoogleOauthResponse struct {
	Ok  bool   `json:"ok"`
	Url string `json:"url"`
}

func NewOauthController(state string) *OauthController {
	return &OauthController{state}
}

func (c *OauthController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.get(w, r)
	default:
		msg := "Method not allowed"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusMethodNotAllowed,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
	}
}

// GET: /google/oauth2
func (c *OauthController) get(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()

	// generate redirect url
	state := c.state
	url := config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	data := GoogleOauthResponse{true, url}

	res, err := json.Marshal(data)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	log.Println("GET: /google/oauth2")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
