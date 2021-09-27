package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yona3/go-auth-sample/controllers"
	controllersGoogle "github.com/yona3/go-auth-sample/controllers/google"
	"github.com/yona3/go-auth-sample/middleware"
)

func Init() {
	r := mux.NewRouter()
	c := middleware.CORS()

	indexController := controllers.NewIndexController()
	googleOauthController := controllersGoogle.NewOauthController()
	googleCallbackController := controllersGoogle.NewCallbackController()

	r.HandleFunc("/", indexController.Index)
	r.HandleFunc("/google/oauth2", googleOauthController.Index)
	r.HandleFunc("/google/callback", googleCallbackController.Index)

	http.Handle("/", c.Handler(r))
}
