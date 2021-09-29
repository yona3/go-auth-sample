package utils

import (
	"context"
	"errors"
	"reflect"
)

type ContextKey string

const (
	UserUUIDContextKey ContextKey = "userUUID"
)

func ContextWithUserUUID(ctx context.Context, userUUID string) context.Context {
	return context.WithValue(ctx, UserUUIDContextKey, userUUID)
}

func GetContextUserUUID(ctx context.Context) (string, error) {
	v := ctx.Value(UserUUIDContextKey)
	if reflect.TypeOf(v).Kind() != reflect.String {
		return "", errors.New("user uuid is not string")
	}
	return v.(string), nil
}
