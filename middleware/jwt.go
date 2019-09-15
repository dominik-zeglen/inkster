package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/core"
)

type UserClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims
}

// UserContextKey defines key holding user data in request context
const UserContextKey = ContextKey("user")

// WithJwt provides jwt token data to request
func WithJwt(next http.Handler, key string) http.Handler {
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
					dataSource := r.Context().Value(UserContextKey).(*core.DataContext)
					user := core.User{}
					user.ID = claims.ID
					dataSource.
						DB().
						Model(&user).
						WherePK().
						First()

					ctx := context.WithValue(r.Context(), UserContextKey, user)
					u := ctx.Value(UserContextKey)
					if u != nil {
						next.ServeHTTP(w, r.WithContext(ctx))
						return
					}
				} else {
					next.ServeHTTP(w, r)
					return
				}
			} else {
				next.ServeHTTP(w, r)
				return
			}
		},
	)
}
