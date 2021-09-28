package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type RedirectOptions struct {
	Type        string // todo: json or redirect
	HttpRequest *http.Request
}

type HandleServerErrorOptions struct {
	Code    int
	Message string
	RedirectOptions
}

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewRedirectOptions(r *http.Request) RedirectOptions {
	return RedirectOptions{
		Type:        "redirect",
		HttpRequest: r,
	}
}

func HandleServerError(w http.ResponseWriter, err error, opts ...HandleServerErrorOptions) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	responseType := "json"
	var httpRequest *http.Request

	if err != nil {
		message = err.Error()
	}

	// set options
	for _, opt := range opts {
		if opt.Code != 0 {
			code = opt.Code
		}

		if opt.Message != "" {
			message = opt.Message
		}

		if opt.Type != "" {
			responseType = opt.Type
		}

		if opt.HttpRequest != nil {
			httpRequest = opt.HttpRequest
		}
	}

	if responseType == "json" {
		// return json
		res := ErrorResponse{false, message}
		json, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(code)
		w.Write([]byte(json))
	} else if responseType == "redirect" {
		// redirect
		if httpRequest == nil {
			log.Println("redirect options is nil")
			panic(err)
		}

		url := "http://localhost:3000/?auth_error=" + message
		http.Redirect(w, httpRequest, url, http.StatusFound)
	} else {
		log.Println("Unknown response type: " + responseType)
		panic("Unknown response type")
	}
}
