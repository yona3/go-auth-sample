package utils

import (
	"net/http"
)

type HandleServerErrorOptions struct {
	Code    int
	Message string
}

func HandleServerError(w http.ResponseWriter, err error, opts ...HandleServerErrorOptions) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if err != nil {
		message = err.Error()
	}

	for _, opt := range opts {
		code = opt.Code
		message = opt.Message
	}

	w.WriteHeader(code)
	w.Write([]byte(message))
}
