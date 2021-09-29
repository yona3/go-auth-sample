package controllersMe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/utils"
)

type MeController struct{}

type TokenResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
}

func NewMeController() *MeController {
	return &MeController{}
}

func (c *MeController) Index(w http.ResponseWriter, r *http.Request) {
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

// GET: /users/me
func (c *MeController) get(w http.ResponseWriter, r *http.Request) {
	// (check access token in middleware!)
	// get user
	// return user
	fmt.Println("GET: /users/me")
}
