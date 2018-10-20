package api

import (
	"context"
	"errors"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	gql "github.com/graph-gophers/graphql-go"
)

type userOperationResult struct {
	errors []userError
	user   *core.User
}

type userOperationResultResolver struct {
	dataSource core.Adapter
	data       userOperationResult
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

type userRemoveResult struct {
	id *gql.ID
}

type userRemoveResultResolver struct {
	data userRemoveResult
}

func (res *userRemoveResultResolver) RemovedObjectID() *gql.ID {
	return res.data.id
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
