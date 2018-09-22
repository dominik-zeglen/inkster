package api

import (
	"errors"

	"github.com/dominik-zeglen/inkster/core"
	// "github.com/globalsign/mgo/bson"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/middleware"
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

func (res *Resolver) User(args UserQueryArgs) (*userResolver, error) {
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

func (res *Resolver) Users() (*[]*userResolver, error) {
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

func (res *Resolver) CreateUser(args UserCreateMutationArgs) (*userOperationResultResolver, error) {
	user := core.User{
		Email: args.Input.Email,
	}
	var pwd string
	if args.Input.Password == nil {
		pwd, _ = user.CreateRandomPassword()
		user.Active = false
	} else {
		user.CreatePassword(*args.Input.Password)
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

func (res *Resolver) RemoveUser(args UserRemoveMutationArgs) (*userRemoveResultResolver, error) {
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return nil, err
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

func (res *Resolver) UpdateUser(args UserUpdateMutationArgs) (*userOperationResultResolver, error) {
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

func (res *Resolver) VerifyToken(args VerifyTokenArgs) bool {
	_, err := jwt.ParseWithClaims(
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
		return false
	}
	return true
}
