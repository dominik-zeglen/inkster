package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mail"
	"github.com/dominik-zeglen/inkster/middleware"
	gql "github.com/graph-gophers/graphql-go"
)

type UserChangePasswordArgs struct {
	ID       gql.ID
	Password string
}

func (res *Resolver) ChangeUserPassword(
	ctx context.Context,
	args UserChangePasswordArgs,
) (bool, error) {
	if !checkPermission(ctx) {
		return false, errNoPermissions
	}
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return false, err
	}

	user := core.User{}
	user.ID = localID
	user.UpdatedAt = res.
		dataSource.
		GetCurrentTime()
	user.CreatePassword(args.Password)

	_, err = res.
		dataSource.
		DB().
		Model(&user).
		Column("updated_at").
		Column("password").
		Column("salt").
		WherePK().
		Update()

	if err != nil {
		return false, err
	}

	return true, nil
}

type ResetUserPasswordArgs struct {
	Password string
	Token    string
}

func (res *Resolver) ResetUserPassword(
	ctx context.Context,
	args ResetUserPasswordArgs,
) (bool, error) {
	appConfig := ctx.Value(middleware.ConfigContextKey).(config.Config)
	tokenObject, err := jwt.ParseWithClaims(
		args.Token,
		&ActionTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
				return nil, errors.New("Invalid signing method")
			}

			key := fmt.Sprintf("%x", appConfig.Server.SecretKey)

			return []byte(key), nil
		},
	)
	if err != nil {
		return false, err
	}

	if claims, ok := tokenObject.Claims.(*ActionTokenClaims); ok {
		user := core.User{}
		user.ID = claims.ID
		user.UpdatedAt = res.
			dataSource.
			GetCurrentTime()
		user.CreatePassword(args.Password)

		_, err = res.
			dataSource.
			DB().
			Model(&user).
			Column("updated_at").
			Column("password").
			Column("salt").
			WherePK().
			Update()
		return true, nil
	}
	return false, nil
}

type SendUserPasswordResetTokenArgs struct {
	Email string
}

func (res *Resolver) SendUserPasswordResetToken(
	ctx context.Context,
	args SendUserPasswordResetTokenArgs,
) (bool, error) {
	appConfig := ctx.Value(middleware.ConfigContextKey).(config.Config)
	website := ctx.Value(middleware.WebsiteContextKey).(core.Website)
	user := core.User{}
	err := res.
		dataSource.
		DB().
		Model(&user).
		Where("email = ?", args.Email).
		First()

	if err != nil {
		return false, nil
	}

	claims := ActionTokenClaims{
		ID:            user.ID,
		AllowedAction: RESET_PASSWORD,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: res.
				dataSource.
				GetCurrentTime().
				Add(time.Hour * 24).
				Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := fmt.Sprintf("%x", appConfig.Server.SecretKey)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return false, err
	}

	err = res.mailer.SendPasswordResetToken(
		args.Email,
		mail.SendPasswordResetTokenTemplateData{
			User:    user,
			Website: website,
			Token:   tokenString,
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}
