package api

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

type pageCreateResult struct {
	userErrors *[]userError
	page       core.Page
}

// Type resolvers
type pageResolver struct {
	dataSource core.Adapter
	data       *core.Page
}
type pageFieldResolver struct {
	dataSource core.Adapter
	data       *core.PageField
}
type pageCreateResultResolver struct {
	dataSource core.Adapter
	data       pageCreateResult
}

func (res *pageCreateResultResolver) Page() *pageResolver {
	return &pageResolver{
		dataSource: res.dataSource,
		data:       &res.data.page,
	}
}

func (res *pageCreateResultResolver) UserErrors() *[]*userErrorResolver {
	var resolverList []*userErrorResolver
	if res.data.userErrors == nil {
		return nil
	}
	userErrors := *res.data.userErrors
	for i := range userErrors {
		resolverList = append(
			resolverList,
			&userErrorResolver{
				data: userErrors[i],
			},
		)
	}
	return &resolverList
}
func (res *pageFieldResolver) Name() string {
	return res.data.Name
}

func (res *pageFieldResolver) Type() string {
	return res.data.Type
}

func (res *pageFieldResolver) Value() *string {
	return &res.data.Value
}

func (res *pageResolver) ID() gql.ID {
	globalID := toGlobalID("page", res.data.ID)
	return gql.ID(globalID)
}

func (res *pageResolver) Name() string {
	return res.data.Name
}

func (res *pageResolver) Fields() *[]*pageFieldResolver {
	var resolverList []*pageFieldResolver
	fields := res.data.Fields
	if fields == nil {
		return nil
	}
	for i := range fields {
		resolverList = append(
			resolverList,
			&pageFieldResolver{
				dataSource: res.dataSource,
				data:       &fields[i],
			},
		)
	}
	return &resolverList
}

func (res *pageResolver) Parent() (*containerResolver, error) {
	parent, err := res.dataSource.GetContainer(res.data.ParentID)
	if err != nil {
		return nil, err
	}
	return &containerResolver{
		dataSource: res.dataSource,
		data:       &parent,
	}, nil
}

type createPageArgs struct {
	Input struct {
		Name     string
		ParentID string
		Fields   *[]*struct {
			Name  string
			Type  string
			Value *string
		}
	}
}

func (res *Resolver) CreatePage(args createPageArgs) (*pageCreateResultResolver, error) {
	if args.Input.Name == "" {
		return &pageCreateResultResolver{
			dataSource: res.dataSource,
			data: pageCreateResult{
				userErrors: &[]userError{
					userError{
						field:   "name",
						message: errNoEmpty("name").Error(),
					},
				},
			},
		}, nil
	}
	if args.Input.ParentID == "" {
		return &pageCreateResultResolver{
			dataSource: res.dataSource,
			data: pageCreateResult{
				userErrors: &[]userError{
					userError{
						field:   "parentId",
						message: errNoEmpty("parentId").Error(),
					},
				},
			},
		}, nil
	}
	localID, err := fromGlobalID("container", args.Input.ParentID)
	if err != nil {
		return nil, err
	}
	page := core.Page{
		Name:     args.Input.Name,
		ParentID: bson.ObjectId(localID),
	}
	if args.Input.Fields != nil {
		fields := *args.Input.Fields
		page.Fields = make([]core.PageField, len(fields))
		for i := range fields {
			page.Fields[i] = core.PageField{
				Name: fields[i].Name,
				Type: fields[i].Type,
			}
		}
	}
	result, err := res.dataSource.AddPage(page)
	if err != nil {
		return nil, err
	}
	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
			page: result,
		},
	}, nil
}

type pageArgs struct {
	ID gql.ID
}

func (res *Resolver) Page(args pageArgs) (*pageResolver, error) {
	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	result, err := res.dataSource.GetPage(localID)
	if err != nil {
		return nil, err
	}
	return &pageResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}
