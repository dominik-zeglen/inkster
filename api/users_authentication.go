package api

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/globalsign/mgo/bson"
)

type loginResult struct {
	token *string
	user  *core.User
}

type loginResultResolver struct {
	dataSource core.Adapter
	data       loginResult
}

func (res *loginResultResolver) Token() *string {
	return res.data.token
}

func (res *loginResultResolver) User() *userResolver {
	if res.data.user != nil {
		return &userResolver{
			data:       res.data.user,
			dataSource: res.dataSource,
		}
	}
	return nil
}

type LoginArgs struct {
	Email    string
	Password string
}

func (res *Resolver) Login(args LoginArgs) (*loginResultResolver, error) {
	user, err := res.dataSource.AuthenticateUser(args.Email, args.Password)
	if err != nil {
		return &loginResultResolver{
			data: loginResult{
				token: nil,
				user:  nil,
			},
			dataSource: res.dataSource,
		}, nil
	}

	claims := middleware.UserClaims{
		Email: user.Email,
		ID:    user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(res.key))
	if err != nil {
		return &loginResultResolver{
			data: loginResult{
				token: nil,
				user:  nil,
			},
			dataSource: res.dataSource,
		}, nil
	}
	return &loginResultResolver{
		data: loginResult{
			token: &tokenString,
			user:  &user,
		},
		dataSource: res.dataSource,
	}, nil
}

type verifyTokenResult struct {
	result bool
	userID *bson.ObjectId
}

type verifyTokenResultResolver struct {
	dataSource core.Adapter
	data       verifyTokenResult
}

func (res *verifyTokenResultResolver) Result() bool {
	return res.data.result
}

func (res *verifyTokenResultResolver) User() (*userResolver, error) {
	if res.data.userID == nil {
		return nil, nil
	}
	user, err := res.dataSource.GetUser(*res.data.userID)
	if err != nil {
		return nil, err
	}
	return &userResolver{
		data:       &user,
		dataSource: res.dataSource,
	}, nil
}

type VerifyTokenArgs struct {
	Token string
}

func (res *Resolver) VerifyToken(args VerifyTokenArgs) *verifyTokenResultResolver {
	tokenObject, err := jwt.ParseWithClaims(
		args.Token,
		&middleware.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
				return nil, errors.New("Invalid signing method")
			}
			return []byte(res.key), nil
		},
	)
	if err != nil {
		return &verifyTokenResultResolver{
			data: verifyTokenResult{
				result: false,
				userID: nil,
			},
			dataSource: res.dataSource,
		}
	}
	claims := tokenObject.Claims.(*middleware.UserClaims)
	id := claims.ID
	return &verifyTokenResultResolver{
		data: verifyTokenResult{
			result: true,
			userID: &id,
		},
		dataSource: res.dataSource,
	}
}
