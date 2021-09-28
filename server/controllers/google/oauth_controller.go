package controllersGoogle

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/utils"
	"golang.org/x/oauth2"
)

type OauthController struct{}

type GoogleOauthResponse struct {
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
		msg := "Method not allowed"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusMethodNotAllowed,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
	}
}

// redirect to AuthCodeURL
func (c *OauthController) get(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()
	// todo: change state
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	data := GoogleOauthResponse{true, url}

	res, err := json.Marshal(data)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
