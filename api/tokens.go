package api

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
)

type Action int

const (
	RESET_PASSWORD Action = iota
)

// JWT claims
type ActionTokenClaims struct {
	ID            bson.ObjectId
	AllowedAction Action
	jwt.StandardClaims
}
