package api

import (
	"context"
	"errors"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/go-pg/pg"
	gql "github.com/graph-gophers/graphql-go"
)

type userOperationResult struct {
	errors []core.ValidationError
	user   *core.User
}

type userOperationResultResolver struct {
	dataSource core.AbstractDataContext
	data       userOperationResult
}

func (res *userOperationResultResolver) Errors() []*inputErrorResolver {
	var resolverList []*inputErrorResolver
	for i := range res.data.errors {
		resolverList = append(
			resolverList,
			&inputErrorResolver{
				err: res.data.errors[i],
			},
		)
	}
	return resolverList
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
	user.CreatedAt = res.
		dataSource.
		GetCurrentTime()
	user.UpdatedAt = res.
		dataSource.
		GetCurrentTime()

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

	validationErrs := user.Validate()
	if len(validationErrs) > 0 {
		return &userOperationResultResolver{
			dataSource: res.dataSource,
			data: userOperationResult{
				errors: validationErrs,
				user:   nil,
			},
		}, nil
	}

	_, err := res.
		dataSource.
		DB().
		Model(&user).
		Insert()

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
			errors: []core.ValidationError{},
			user:   &user,
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
				errors: validationErrors,
				user:   cleanedUser,
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
			errors: []core.ValidationError{},
			user:   &user,
		},
	}, nil
}
