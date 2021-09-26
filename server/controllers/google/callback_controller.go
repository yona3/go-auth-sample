package controllersGoogle

import (
	"context"
	"encoding/json"
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

type ResponseData struct {
	UserId      string `json:"user_id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
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

	// todo: set refresh_token

	data := ResponseData{
		UserId:      info.UserId,
		Email:       info.Email,
		AccessToken: tok.AccessToken,
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println(string(b))

	// return json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
