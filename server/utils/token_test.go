package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestSetRefreshToken(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		uuid uuid.UUID
	}

	recorder := httptest.NewRecorder()

	cases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{recorder, uuid.New()}, false},
		{"fail", args{recorder, uuid.Nil}, true},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetRefreshToken(tt.args.w, tt.args.uuid); (err != nil) != tt.wantErr {
				t.Errorf("SetRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.name == "pass" {
				// check cookie value
				c := tt.args.w.Header().Get("Set-Cookie")
				if c == "" {
					t.Error("SetRefreshToken() cookie not set")
				}

				// check is jwt
				token := strings.Split(strings.Split(c, ";")[0], "=")[1]
				if !strings.HasPrefix(token, "eyJ") {
					t.Error("token is not jwt")
				}

				// todo: check other values...
			}
		})
	}
}

func TestRevokeRefreshToken(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	// new requests
	passRequest, err := http.NewRequest(http.MethodDelete, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	failRequest := httptest.NewRequest(http.MethodDelete, "/", nil)

	recorder := httptest.NewRecorder()

	// set cookie
	passRequest.AddCookie(&http.Cookie{
		Name:  "token",
		Value: "test",
	})

	cases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{recorder, passRequest}, false},
		{"fail", args{recorder, failRequest}, true},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if err := RevokeRefreshToken(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("RevokeRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.name == "pass" {
				// check cookie
				c := tt.args.w.Header().Get("Set-Cookie")
				if c == "" {
					t.Error("RevokeRefreshToken() cookie not set")
				}

				values := strings.Split(c, ";")

				// check is empty
				token := strings.Split(values[0], "=")[1]
				if token != "" {
					t.Error("token is not empty")
				}

				// check max age is 0
				maxAge := strings.Split(values[1], "=")[1]
				if maxAge != "0" {
					t.Error("max age is not 0")
				}
			}
		})
	}
}
