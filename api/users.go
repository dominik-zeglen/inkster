package api

import (
	"context"
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

// Type resolver
type userResolver struct {
	dataSource core.Adapter
	data       *core.User
}
type userOperationResult struct {
	errors []userError
	user   *core.User
}
type userOperationResultResolver struct {
	dataSource core.Adapter
	data       userOperationResult
}
type userRemoveResult struct {
	id *gql.ID
}
type userRemoveResultResolver struct {
	data userRemoveResult
}
type loginResult struct {
	token *string
	user  *core.User
}
type loginResultResolver struct {
	dataSource core.Adapter
	data       loginResult
}
type verifyTokenResult struct {
	result bool
	userID *bson.ObjectId
}
type verifyTokenResultResolver struct {
	dataSource core.Adapter
	data       verifyTokenResult
}

func (res *userOperationResultResolver) Errors() []*userErrorResolver {
	var resolverList []*userErrorResolver
	for i := range res.data.errors {
		resolverList = append(
			resolverList,
			&userErrorResolver{
				data: res.data.errors[i],
			},
		)
	}
	return resolverList
}
func (res *userOperationResultResolver) User() *userResolver {
	return &userResolver{
		dataSource: res.dataSource,
		data:       res.data.user,
	}
}

func (res *userRemoveResultResolver) RemovedObjectID() *gql.ID {
	return res.data.id
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

func (res *userResolver) ID() gql.ID {
	globalID := toGlobalID("user", res.data.ID)
	return gql.ID(globalID)
}

func (res *userResolver) CreatedAt() string {
	return res.data.CreatedAt
}

func (res *userResolver) UpdatedAt() string {
	return res.data.UpdatedAt
}

func (res *userResolver) Email() string {
	return res.data.Email
}

func (res *userResolver) IsActive() bool {
	return res.data.Active
}

type UserQueryArgs struct {
	ID gql.ID
}

func (res *Resolver) User(
	ctx context.Context,
	args UserQueryArgs,
) (*userResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return nil, err
	}
	result, err := res.dataSource.GetUser(localID)
	if err != nil {
		return nil, err
	}
	return &userResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}

func (res *Resolver) Users(ctx context.Context) (*[]*userResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	var resolverList []*userResolver
	result, err := res.dataSource.GetUserList()
	if err != nil {
		return nil, err
	}
	for index := range result {
		resolverList = append(
			resolverList,
			&userResolver{
				dataSource: res.dataSource,
				data:       &result[index],
			},
		)
	}
	return &resolverList, nil
}

type UserCreateInput struct {
	Email    string
	Password *string
}
type UserCreateMutationArgs struct {
	Input          UserCreateInput
	SendInvitation *bool
}

func (res *Resolver) CreateUser(
	ctx context.Context,
	args UserCreateMutationArgs,
) (*userOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	user := core.User{
		Email: args.Input.Email,
	}
	var pwd string
	if args.Input.Password == nil {
		pwd, _ = user.CreateRandomPassword()
		user.Active = false
	} else {
		err := user.CreatePassword(*args.Input.Password)
		if err != nil {
			return nil, err
		}
		user.Active = true
	}
	result, err := res.dataSource.AddUser(user)
	if err != nil {
		return nil, err
	}
	if args.SendInvitation != nil {
		sendInvitation := *args.SendInvitation
		if sendInvitation {
			err = res.mailer.Send(user.Email, "Inkster password", pwd)
			if err != nil {
				return nil, err
			}
		}
	}
	return &userOperationResultResolver{
		dataSource: res.dataSource,
		data: userOperationResult{
			errors: []userError{},
			user:   &result,
		},
	}, nil
}

type UserRemoveMutationArgs struct {
	ID gql.ID
}

func (res *Resolver) RemoveUser(
	ctx context.Context,
	args UserRemoveMutationArgs,
) (*userRemoveResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return nil, err
	}

	if user, ok := ctx.Value("user").(*middleware.UserClaims); ok {
		if user.ID == localID {
			return nil, errors.New("User cannot remove himself")
		}
	}

	err = res.dataSource.RemoveUser(localID)
	if err != nil {
		return nil, err
	}
	return &userRemoveResultResolver{
		data: userRemoveResult{
			id: &args.ID,
		},
	}, nil
}

type UserUpdateInput struct {
	IsActive *bool
	Email    *string
}
type UserUpdateMutationArgs struct {
	ID    gql.ID
	Input UserUpdateInput
}

func (res *Resolver) UpdateUser(
	ctx context.Context,
	args UserUpdateMutationArgs,
) (*userOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return nil, err
	}
	input := core.UserInput{
		Active: args.Input.IsActive,
		Email:  args.Input.Email,
	}
	result, err := res.dataSource.UpdateUser(localID, input)
	if err != nil {
		return nil, err
	}
	return &userOperationResultResolver{
		dataSource: res.dataSource,
		data: userOperationResult{
			errors: []userError{},
			user:   &result,
		},
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
	_, err = res.dataSource.UpdateUser(localID, core.UserInput{
		Password: &args.Password,
	})
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

			user, err := res.dataSource.GetUser(bson.ObjectId(claims.ID))
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
		_, err = res.dataSource.UpdateUser(claims.ID, core.UserInput{
			Password: &args.Password,
		})
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
	user, err := res.dataSource.GetUserByEmail(args.Email)
	if err != nil {
		return true, nil
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
