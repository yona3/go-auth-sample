package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yona3/go-auth-sample/controllers"
	controllersGoogle "github.com/yona3/go-auth-sample/controllers/google"
	controllersRefreshToken "github.com/yona3/go-auth-sample/controllers/refresh-token"
	controllersToken "github.com/yona3/go-auth-sample/controllers/token"
	controllersMe "github.com/yona3/go-auth-sample/controllers/user/me"
	"github.com/yona3/go-auth-sample/middlewares"
	"github.com/yona3/go-auth-sample/utils"
)

func Init() {
	r := mux.NewRouter()
	c := NewCors()
	state := utils.GenerateRandomString(32)

	indexController := controllers.NewIndexController()
	googleOauthController := controllersGoogle.NewOauthController(state)
	googleCallbackController := controllersGoogle.NewCallbackController(state)
	tokenController := controllersToken.NewTokenController()
	refreshTokenController := controllersRefreshToken.NewRefreshTokenController()
	meController := controllersMe.NewMeController()

	r.HandleFunc("/", indexController.Index)
	r.HandleFunc("/google/oauth2", googleOauthController.Index)
	r.HandleFunc("/google/callback", googleCallbackController.Index)
	r.HandleFunc("/token", tokenController.Index)
	r.HandleFunc("/refresh_token", middlewares.AuthToken(refreshTokenController.Index))
	r.HandleFunc("/users/me", middlewares.AuthToken(meController.Index))

	http.Handle("/", c.Handler(r))
}
