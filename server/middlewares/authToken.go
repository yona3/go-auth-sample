package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/yona3/go-auth-sample/utils"
)

func AuthToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get access token from header
		v := r.Header.Get("Authorization")
		if v == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(v, " ")[1]
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// verify jwt
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// check algoritm
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

		// set claims value to request
		ctx := r.Context()
		ctx = utils.ContextWithUserUUID(ctx, userUUID)

		log.Println("[PASS]: AuthToken middleware")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
