package api

import (
	"context"
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/core"
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
	tokenObject, err := jwt.ParseWithClaims(
		args.Token,
		&ActionTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
				return nil, errors.New("Invalid signing method")
			}

			claims, ok := token.Claims.(*ActionTokenClaims)
			if !ok {
				return nil, errors.New("Invalid token claims")
			}

			user := core.User{}
			user.ID = claims.ID

			err := res.
				dataSource.
				DB().
				Model(&user).
				WherePK().
				Select()

			if err != nil {
				return nil, err
			}

			key := fmt.Sprintf("%x", user.Password)

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
	user := core.User{}
	err := dataSource.
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
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := fmt.Sprintf("%x", user.Password)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return false, err
	}

	err = res.mailer.Send(args.Email, "Inkster reset password", tokenString)
	if err != nil {
		return false, err
	}

	return true, nil
}
