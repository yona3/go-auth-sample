package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yona3/go-auth-sample/controllers"
)

func Init() {
	r := mux.NewRouter()
	indexController := controllers.NewIndexController()

	r.HandleFunc("/", indexController.Index)
	http.Handle("/", r)
}
