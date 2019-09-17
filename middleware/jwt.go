package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/core"
)

// UserClaims holds all token data
type UserClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// UserContextKey defines key holding user data in request context
const UserContextKey = ContextKey("user")

// WithJwt provides jwt token data to request
func WithJwt(
	next http.Handler,
	key string,
	dataSource core.DataContext,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			headerContent := r.Header.Get("Authorization")
			if headerContent != "" && headerContent != "null" {
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
					user := core.User{}
					user.ID = claims.ID
					dataSource.
						DB().
						Model(&user).
						WherePK().
						First()

					ctx := context.WithValue(r.Context(), UserContextKey, &user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					next.ServeHTTP(w, r)
				}
			} else {
				next.ServeHTTP(w, r)
			}

			return
		},
	)
}
