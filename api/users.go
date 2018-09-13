package api

import (
	"github.com/dominik-zeglen/inkster/core"
	// "github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

// type userUpdateResult struct {
// 	userErrors *[]userError
// 	userID     bson.ObjectId
// }
// type userRemoveResult struct {
// 	userErrors *[]userError
// 	userID     bson.ObjectId
// }

// Type resolver
type userResolver struct {
	dataSource core.Adapter
	data       *core.User
}
type userCreateResult struct {
	errors []userError
	user   *core.User
}
type userCreateResultResolver struct {
	dataSource core.Adapter
	data       userCreateResult
}
type userRemoveResult struct {
	id *gql.ID
}
type userRemoveResultResolver struct {
	data userRemoveResult
}

// type userUpdateResultResolver struct {
// 	dataSource core.ADapter
// 	data       userUpdateResult
// }

// func (res *userUpdateResultResolver) User() (*userResolver, error) {
// 	result, err := res.dataSource.GetUser(res.data.userID)
// 	if err != nil {
// 		dyppy
// 		return err
// 	}
// 	return &userResolver{
// 		dataSource: res.dataSource,
// 		data:       &result,
// 	}
// }
//
// func (res *userUpdateResultResolver) UserErrors() *[]*userErrorResolver {
// 	var resolverList []*userErrorResolver
// 	if res.data.userErrors == nil {
// 		return nil
// 	}
// 	userErrors := *res.data.userErrors
// 	for i := range userErrors {
// 		resolverList = append(
// 			resolverList,
// 			&userErrorResolver{
// 				data: userErrors[i],
// 			},
// 		)
// 	}
// 	return &resolverList
// }
// func (res *userRemoveResultResolver) RemovedObjectID() *gql.ID {
// 	id := gql.ID(toGlobalID("user", res.data.userID))
// 	return &id
// }

func (res *userCreateResultResolver) Errors() []*userErrorResolver {
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
func (res *userCreateResultResolver) User() *userResolver {
	return &userResolver{
		data: res.data.user,
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
	Input UserCreateInput
}

func (res *Resolver) CreateUser(args UserCreateMutationArgs) (*userCreateResultResolver, error) {
	user := core.User{
		Email: args.Input.Email,
	}
	if args.Input.Password == nil {
		user.CreateRandomPassword()
		user.Active = false
	} else {
		user.CreatePassword(*args.Input.Password)
		user.Active = true
	}
	result, err := res.dataSource.AddUser(user)
	if err != nil {
		return nil, err
	}
	return &userCreateResultResolver{
		dataSource: res.dataSource,
		data: userCreateResult{
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
