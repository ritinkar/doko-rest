package controllers

import (
	"context"
	"doko-rest/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type key string

const (
	tokenKey key = "Token"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(os.Getenv("JWT_SECRET")), nil
				})
				if error != nil {
					utils.Respond(w, utils.Message(false, error.Error()))
					return
				}
				if token.Valid {
					ctx := req.Context()
					ctx = context.WithValue(ctx, tokenKey, token)
					req = req.WithContext(ctx)
					next(w, req)
				} else {
				}
			}
		} else {
			utils.Respond(w, utils.Message(false, "An authorization header is required"))

		}
	})
}
