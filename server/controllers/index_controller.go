package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexController struct{}

type IndexResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.get(w, r)
	default:
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

// GET: /
func (c *IndexController) get(w http.ResponseWriter, r *http.Request) {
	data := IndexResponse{true, "Hello, world!"}

	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
