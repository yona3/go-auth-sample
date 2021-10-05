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

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				w:    httptest.NewRecorder(),
				uuid: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				w:    httptest.NewRecorder(),
				uuid: uuid.Nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
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

	// new pass request
	passRequest, err := http.NewRequest(http.MethodDelete, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// set cookie
	passRequest.AddCookie(&http.Cookie{
		Name:  "token",
		Value: "test",
	})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				w: httptest.NewRecorder(),
				r: passRequest,
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/", nil),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
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
