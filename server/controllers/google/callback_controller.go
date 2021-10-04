package controllersGoogle

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/database"
	"github.com/yona3/go-auth-sample/ent"
	"github.com/yona3/go-auth-sample/ent/user"
	"github.com/yona3/go-auth-sample/utils"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	userData          *ent.User
	redirectOptions   utils.RedirectOptions
	defaultOptions    utils.HandleServerErrorOptions
	handleServerError func(w http.ResponseWriter, err error, opts ...utils.HandleServerErrorOptions)
)

type CallbackController struct {
	state string
}

type GoogleCallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func NewCallbackController(state string) *CallbackController {
	return &CallbackController{state}
}

func (c *CallbackController) Index(w http.ResponseWriter, r *http.Request) {
	// init
	redirectOptions = utils.NewRedirectOptions(r)
	defaultOptions = utils.HandleServerErrorOptions{
		RedirectOptions: redirectOptions,
	}
	handleServerError = utils.NewHandleServerErrorWithDefaultOptions(defaultOptions)

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
		handleServerError(w, nil, opts)
	}
}

// GET: /google/callback
func (c *CallbackController) get(w http.ResponseWriter, r *http.Request) {
	req := GoogleCallbackRequest{}

	// get form values
	req.Code = r.FormValue("code")
	req.State = r.FormValue("state")

	config := GetConfig()
	state := c.state

	// check state is valid
	if req.State != state {
		msg := "Invalid state"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		handleServerError(w, nil, opts)
		return
	}

	ctx := context.Background()

	// get token
	tok, err := config.Exchange(ctx, req.Code)
	if err != nil {
		handleServerError(w, err)
		return
	}

	// check token is valid
	if !tok.Valid() {
		msg := "invalid token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		handleServerError(w, err, opts)
		return
	}

	// check refresh token is empty
	if tok.RefreshToken == "" {
		msg := "refresh token is empty"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		handleServerError(w, err, opts)
		return
	}

	opts := option.WithHTTPClient(config.Client(ctx, tok))

	s, err := v2.NewService(ctx, opts)
	if err != nil {
		handleServerError(w, err)
		return
	}

	// get user info from service
	info, err := s.Userinfo.V2.Me.Get().Do()
	if err != nil {
		handleServerError(w, err)
		return
	}

	// get user info from database
	db := database.GetClient()
	u, err := db.User.Query().Where(user.Email(info.Email)).Only(ctx)
	if err != nil {
		fmt.Println("user not found")

		// create user
		new, err := db.User.Create().SetEmail(info.Email).SetName(info.Name).SetSigninWith("google").Save(ctx)
		if err != nil {
			handleServerError(w, err)
			return
		}

		userData = new
		log.Printf("user created. (uuid: %v)\n", userData.UUID)
	} else {
		// check signin method
		if u.SigninWith != "google" {
			msg := "this email has already been used"
			log.Println(msg)

			opts := utils.HandleServerErrorOptions{
				Code:    http.StatusBadRequest,
				Message: msg,
			}
			handleServerError(w, nil, opts)
			return
		}

		userData = u
		log.Printf("user found. (uuid: %v)\n", userData.UUID)
	}

	if userData == nil {
		handleServerError(w, err)
		return
	}

	// set refresh token to cookie
	if utils.SetRefreshToken(w, userData.UUID); err != nil {
		msg := "failed to generate token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		handleServerError(w, err, opts)
		return
	}

	log.Println("GET: /google/callback")

	url := "http://localhost:3000"
	http.Redirect(w, r, url, http.StatusFound)
}
