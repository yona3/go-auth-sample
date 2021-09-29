package controllersUser

import (
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/utils"
)

type UserController struct{}

type TokenResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
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
