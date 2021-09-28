package controllersGoogle

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/utils"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type CallbackController struct{}

type GoogleCallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func NewCallbackController() *CallbackController {
	return &CallbackController{}
}

func (c *CallbackController) Index(w http.ResponseWriter, r *http.Request) {
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

func (c *CallbackController) get(w http.ResponseWriter, r *http.Request) {
	req := GoogleCallbackRequest{}

	// get form values
	req.Code = r.FormValue("code")
	req.State = r.FormValue("state")

	config := GetConfig()
	ctx := context.Background()

	// todo: check state

	tok, err := config.Exchange(ctx, req.Code)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	if !tok.Valid() {
		msg := "invalid token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, err, opts)
		return
	}

	// check refresh token is empty
	if tok.RefreshToken == "" {
		msg := "refresh token is empty"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, err, opts)
		return
	}

	opts := option.WithHTTPClient(config.Client(ctx, tok))

	s, err := v2.NewService(ctx, opts)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	// get user info
	info, err := s.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		log.Println(err.Error())
		utils.HandleServerError(w, err)
		return
	}

	fmt.Printf("user is logged in. (email: %v)\n", info.Email)

	fmt.Printf("access token: %v\n", tok.AccessToken)
	fmt.Printf("refresh token: %v\n", tok.RefreshToken)

	// todo: set refresh_token

	url := "http://localhost:3000"
	http.Redirect(w, r, url, http.StatusFound)
}
