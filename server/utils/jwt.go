package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTClaims struct {
	UserUUID uuid.UUID `json:"user_uuid"`
	jwt.StandardClaims
}

func NewJWTClaims(userUUID uuid.UUID, exp int64) *JWTClaims {
	return &JWTClaims{
		UserUUID: userUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
}
