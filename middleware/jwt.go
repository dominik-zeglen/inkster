package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func WithJwt(next http.Handler, key string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			headerContent := r.Header.Get("Authorization")
			if headerContent != "" {
				tokenString := strings.Split(headerContent, " ")[1]
				token, err := jwt.ParseWithClaims(
					tokenString,
					&UserClaims{},
					func(token *jwt.Token) (interface{}, error) {
						if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
							return nil, errors.New("Invalid signing method")
						}
						return []byte(key), nil
					},
				)

				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				if claims, valid := token.Claims.(*UserClaims); valid && token.Valid {
					ctx := context.WithValue(r.Context(), "user", claims)
					u := ctx.Value("user")
					if u != nil {
						next.ServeHTTP(w, r.WithContext(ctx))
					}
				} else {
					next.ServeHTTP(w, r)
				}
			} else {
				next.ServeHTTP(w, r)
			}
		},
	)
}
