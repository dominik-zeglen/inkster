package api

import (
	"context"
	"errors"

	"github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mail"
	"github.com/dominik-zeglen/inkster/middleware"
	gql "github.com/graph-gophers/graphql-go"
)

type userOperationResult struct {
	user             *core.User
	validationErrors []core.ValidationError
}

type userOperationResultResolver struct {
	dataSource core.AbstractDataContext
	data       userOperationResult
}

func (res *userOperationResultResolver) Errors() []inputErrorResolver {
	return createInputErrorResolvers(res.data.validationErrors)
}

func (res *userOperationResultResolver) User() *userResolver {
	if res.data.user == nil {
		return nil
	}
	return &userResolver{
		dataSource: res.dataSource,
		data:       res.data.user,
	}
}

type UserCreateMutationArgs struct {
	Input          core.UserCreateInput
	SendInvitation *bool
}

func (res *Resolver) CreateUser(
	ctx context.Context,
	args UserCreateMutationArgs,
) (*userOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	appConfig := ctx.Value(middleware.ConfigContextKey).(config.Config)
	website := ctx.Value(middleware.WebsiteContextKey).(core.Website)

	user, validationErrors, err := core.CreateUser(
		args.Input,
		res.dataSource,
	)

	if err != nil || len(validationErrors) != 0 {
		return &userOperationResultResolver{
			dataSource: res.dataSource,
			data: userOperationResult{
				validationErrors: validationErrors,
				user:             nil,
			},
		}, err
	}

	token, err := createPasswordResetToken(
		user.ID,
		res.dataSource.GetCurrentTime(),
		appConfig.Server.SecretKey,
	)

	err = res.mailer.SendUserInvitation(
		user.Email,
		mail.SendUserInvitationTemplateData{
			User:    *user,
			Website: website,
			Token:   token,
		},
	)

	return &userOperationResultResolver{
		dataSource: res.dataSource,
		data: userOperationResult{
			validationErrors: []core.ValidationError{},
			user:             user,
		},
	}, err
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

	user := core.User{}
	user.ID = localID

	err = res.
		dataSource.
		DB().
		Model(&user).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	if args.Input.Email != nil {
		user.Email = *args.Input.Email
	}
	if args.Input.IsActive != nil {
		user.Active = *args.Input.IsActive
	}

	updatedUser, validationErrors, err := core.UpdateUser(
		user,
		res.dataSource,
	)

	return &userOperationResultResolver{
		data: userOperationResult{
			user:             updatedUser,
			validationErrors: validationErrors,
		},
		dataSource: res.dataSource,
	}, err
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

type userRemoveMutationArgs struct {
	ID gql.ID
}

func (res *Resolver) RemoveUser(
	ctx context.Context,
	args userRemoveMutationArgs,
) (*userRemoveResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return nil, err
	}

	if user, ok := ctx.Value(middleware.UserContextKey).(*core.User); ok {
		if user.ID == localID {
			return nil, errors.New("User cannot remove himself")
		}
	}

	err = core.RemoveUser(localID, res.dataSource)
	if err != nil {
		return nil, err
	}

	return &userRemoveResultResolver{
		data: userRemoveResult{
			id: &args.ID,
		},
	}, nil
}
