package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yona3/go-auth-sample/utils"
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
		msg := "Method not allowed"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusMethodNotAllowed,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
	}
}

// GET: /
func (c *IndexController) get(w http.ResponseWriter, r *http.Request) {
	data := IndexResponse{true, "Hello, world!"}

	res, err := json.Marshal(data)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
