package controllers

import (
	"log"
	"net/http"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.getIndex(w, r)
	default:
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

// GET: /
func (c *IndexController) getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
