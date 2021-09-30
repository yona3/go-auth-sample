package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func SetRefreshToken(w http.ResponseWriter, uuid uuid.UUID) error {
	// generate refresh token
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := NewJWTClaims(uuid, expiresAt)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStrnig, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	// set refresh token to cookie
	c := &http.Cookie{
		Path:     "/",
		Name:     "token",
		Value:    tokenStrnig,
		MaxAge:   60 * 60 * 24 * 7, // 7 days
		HttpOnly: true,
		Secure:   false, // !change to true when deploy to production
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, c)

	return nil
}

func RevokeRefreshToken(w http.ResponseWriter, r *http.Request) error {
	c, err := r.Cookie("token")
	if err != nil {
		return err
	}

	c.MaxAge = -1
	http.SetCookie(w, c)

	return nil
}
