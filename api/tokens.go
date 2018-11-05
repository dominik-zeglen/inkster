package api

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Action int

const (
	RESET_PASSWORD Action = iota
)

// JWT claims
type ActionTokenClaims struct {
	ID            string
	AllowedAction Action
	jwt.StandardClaims
}
