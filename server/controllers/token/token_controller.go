package controllersToken

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/yona3/go-auth-sample/database"
	"github.com/yona3/go-auth-sample/ent/user"
	"github.com/yona3/go-auth-sample/utils"
)

type TokenController struct{}

type TokenResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
}

func NewTokenController() *TokenController {
	return &TokenController{}
}

func (c *TokenController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.post(w, r)
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

func (c *TokenController) post(w http.ResponseWriter, r *http.Request) {
	// get cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		msg := "Cookie not found"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	v := cookie.Value

	// verify jwt
	token, err := jwt.Parse(v, func(token *jwt.Token) (interface{}, error) {
		if ok := token.Method.Alg() == "HS256"; !ok {
			return nil, errors.New(fmt.Sprint("Unexpected signing method: ", token.Method.Alg()))
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)

		opts := utils.HandleServerErrorOptions{
			Code: http.StatusBadRequest,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// get jwt claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		msg := "Invalid token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// get claims value
	userUUID := claims["user_uuid"].(string)
	exp := claims["exp"].(float64)

	// parse uuid from string
	id, err := uuid.Parse(userUUID)
	if err != nil {
		msg := "UUID is invalid"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// search user by uuid
	ctx := context.Background()
	db := database.GetClient()
	u, err := db.User.Query().Where(user.UUID(id)).Only(ctx)
	if err != nil {
		msg := "User not found"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// check logged out at
	// if the token was generated before logout, it will be invalid.
	if u.LoggedOutAt.Unix() >= int64(exp) {
		msg := "User is logged out"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// generate access_token
	newExp := time.Now().Add(time.Minute * 15).Unix() // 15 minutes
	newClaims := utils.NewJWTClaims(u.UUID, newExp)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		msg := "Failed to generate access token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// refresh refresh token
	if err := utils.SetRefreshToken(w, u.UUID); err != nil {
		msg := "Failed to set refresh token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// return access_token
	data := TokenResponse{
		Ok:          true,
		AccessToken: tokenString,
	}

	res, err := json.Marshal(data)
	if err != nil {
		msg := "failed to json marshal"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
	}

	log.Println("POST: /token")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
