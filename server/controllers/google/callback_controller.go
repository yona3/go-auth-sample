package controllersGoogle

import (
	"context"
	"fmt"
	"log"
	"net/http"

	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type CallbackController struct{}

type CallbackRequest struct {
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
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func (c *CallbackController) get(w http.ResponseWriter, r *http.Request) {
	req := CallbackRequest{}

	// get form values
	req.Code = r.FormValue("code")
	req.State = r.FormValue("state")

	config := GetConfig()
	ctx := context.Background()

	// todo: check state

	tok, err := config.Exchange(ctx, req.Code)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !tok.Valid() {
		log.Println("valid token")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// check refresh token is empty
	if tok.RefreshToken == "" {
		log.Println("refresh token is empty")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	opts := option.WithHTTPClient(config.Client(ctx, tok))

	s, err := v2.NewService(ctx, opts)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// get user info
	info, err := s.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("user is logged in. (email: %v)\n", info.Email)

	fmt.Printf("access token: %v\n", tok.AccessToken)
	fmt.Printf("refresh token: %v\n", tok.RefreshToken)

	// todo: set refresh_token

	url := "http://localhost:3000"
	http.Redirect(w, r, url, http.StatusFound)
}
