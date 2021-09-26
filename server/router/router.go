package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yona3/go-auth-sample/controllers"
	controllersGoogle "github.com/yona3/go-auth-sample/controllers/google"
)

func Init() {
	r := mux.NewRouter()

	indexController := controllers.NewIndexController()
	googleOauthController := controllersGoogle.NewOauthController()
	googleCallbackController := controllersGoogle.NewCallbackController()

	r.HandleFunc("/", indexController.Index)
	r.HandleFunc("/google/oauth2", googleOauthController.Get)
	r.HandleFunc("/google/callback", googleCallbackController.Get)

	http.Handle("/", r)
}
