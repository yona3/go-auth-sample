package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yona3/go-auth-sample/controllers"
	controllersGoogle "github.com/yona3/go-auth-sample/controllers/google"
	controllersToken "github.com/yona3/go-auth-sample/controllers/token"
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

	r.HandleFunc("/", indexController.Index)
	r.HandleFunc("/token", tokenController.Index)
	r.HandleFunc("/google/oauth2", googleOauthController.Index)
	r.HandleFunc("/google/callback", googleCallbackController.Index)

	http.Handle("/", c.Handler(r))
}
