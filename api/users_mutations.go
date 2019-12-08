package api

import (
	"context"
	"errors"

	"github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mail"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/go-pg/pg"
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
	Email string
}
type UserCreateMutationArgs struct {
	Input UserCreateInput
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
	user := core.User{
		Email: args.Input.Email,
	}

	user.CreatedAt = res.
		dataSource.
		GetCurrentTime()
	user.UpdatedAt = res.
		dataSource.
		GetCurrentTime()

	_, err := user.CreateRandomPassword()

	validationErrs := user.Validate()
	if len(validationErrs) > 0 {
		return &userOperationResultResolver{
			dataSource: res.dataSource,
			data: userOperationResult{
				validationErrors: validationErrs,
				user:             nil,
			},
		}, err
	}

	_, err = res.
		dataSource.
		DB().
		Model(&user).
		Insert()

	if err != nil {
		return &userOperationResultResolver{
			dataSource: res.dataSource,
			data: userOperationResult{
				validationErrors: []core.ValidationError{},
				user:             &user,
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
			User:    user,
			Website: website,
			Token:   token,
		},
	)

	return &userOperationResultResolver{
		dataSource: res.dataSource,
		data: userOperationResult{
			validationErrors: []core.ValidationError{},
			user:             &user,
		},
	}, err
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

	if user, ok := ctx.Value(middleware.UserContextKey).(*core.User); ok {
		if user.ID == localID {
			return nil, errors.New("User cannot remove himself")
		}
	}

	user := core.User{}
	user.ID = localID
	_, err = res.
		dataSource.
		DB().
		Model(&user).
		WherePK().
		Delete()
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
	Email    *string `validate:"omitempty,email"`
}
type UserUpdateMutationArgs struct {
	ID    gql.ID
	Input UserUpdateInput `validate:"dive"`
}

func (args UserUpdateMutationArgs) validate(
	dataSource core.AbstractDataContext,
	userID int,
) (
	[]core.ValidationError,
	*core.User,
	error,
) {
	errors := []core.ValidationError{}
	errors = append(errors, core.ValidateModel(args)...)

	user := core.User{}
	user.ID = userID

	err := dataSource.
		DB().
		Model(&user).
		WherePK().
		Select()

	if err != nil {
		return errors, nil, err
	}

	if args.Input.Email != nil && *args.Input.Email != user.Email {
		user := core.User{}
		err := dataSource.
			DB().
			Model(&user).
			Where("email = ?", *args.Input.Email).
			First()

		if err != nil {
			if err != pg.ErrNoRows {
				return errors, nil, err
			}
		} else {
			errors = append(errors, core.ValidationError{
				Code:  core.ErrNotUnique,
				Field: "email",
				Param: args.Input.Email,
			})
		}
	}

	return errors, &user, nil
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

	validationErrors, cleanedUser, err := args.validate(res.dataSource, localID)
	if err != nil {
		return nil, err
	}
	if len(validationErrors) > 0 {
		return &userOperationResultResolver{
			dataSource: res.dataSource,
			data: userOperationResult{
				validationErrors: validationErrors,
				user:             cleanedUser,
			},
		}, nil
	}

	user := *cleanedUser
	user.ID = localID
	user.UpdatedAt = res.
		dataSource.
		GetCurrentTime()

	query := res.
		dataSource.
		DB().
		Model(&user).
		Column("updated_at")

	if args.Input.Email != nil {
		user.Email = *args.Input.Email
		query = query.Column("email")
	}
	if args.Input.IsActive != nil {
		user.Active = *args.Input.IsActive
		query = query.Column("active")
	}

	_, err = query.
		WherePK().
		Update()

	if err != nil {
		return nil, err
	}
	return &userOperationResultResolver{
		dataSource: res.dataSource,
		data: userOperationResult{
			validationErrors: []core.ValidationError{},
			user:             &user,
		},
	}, nil
}
