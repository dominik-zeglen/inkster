package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

type templateUpdateResult struct {
	userErrors *[]userError
	templateID bson.ObjectId
}
type templateRemoveResult struct {
	userErrors *[]userError
	templateID bson.ObjectId
}

// Type resolvers
type templateResolver struct {
	dataSource core.Adapter
	data       *core.Template
}
type templateFieldResolver struct {
	dataSource core.Adapter
	data       *core.TemplateField
}
type templateUpdateResultResolver struct {
	dataSource core.Adapter
	data       templateUpdateResult
}
type templateRemoveResultResolver struct {
	dataSource core.Adapter
	data       templateRemoveResult
}

func (res *templateUpdateResultResolver) Template() *templateResolver {
	result, err := res.dataSource.GetTemplate(res.data.templateID)
	if err != nil {
		return nil
	}
	return &templateResolver{
		dataSource: res.dataSource,
		data:       &result,
	}
}

func (res *templateUpdateResultResolver) UserErrors() *[]*userErrorResolver {
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
func (res *templateRemoveResultResolver) RemovedObjectID() *gql.ID {
	id := gql.ID(toGlobalID("template", res.data.templateID))
	return &id
}

func (res *templateRemoveResultResolver) UserErrors() *[]*userErrorResolver {
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

func (res *templateFieldResolver) Name() string {
	return res.data.Name
}

func (res *templateFieldResolver) Type() string {
	return res.data.Type
}

func (res *templateResolver) ID() gql.ID {
	globalID := toGlobalID("template", res.data.ID)
	return gql.ID(globalID)
}

func (res *templateResolver) Name() string {
	return res.data.Name
}

func (res *templateResolver) Fields() *[]*templateFieldResolver {
	var resolverList []*templateFieldResolver
	fields := res.data.Fields
	if fields == nil {
		return nil
	}
	for i := range fields {
		resolverList = append(
			resolverList,
			&templateFieldResolver{
				dataSource: res.dataSource,
				data:       &fields[i],
			},
		)
	}
	return &resolverList
}

type createTemplateArgs struct {
	Input struct {
		Name   string
		Fields *[]struct {
			Name string
			Type string
		}
	}
}

func (res *Resolver) CreateTemplate(args createTemplateArgs) (
	*templateResolver, error,
) {
	input := &args.Input
	template := core.Template{
		Name: input.Name,
	}
	if input.Fields != nil {
		fields := *input.Fields
		template.Fields = make([]core.TemplateField, len(fields))
		for i := range fields {
			template.Fields[i] = core.TemplateField{
				Name: fields[i].Name,
				Type: fields[i].Type,
			}
		}
	}
	result, err := res.dataSource.AddTemplate(template)
	if err != nil {
		return nil, err
	}
	return &templateResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}

func (res *Resolver) Template(args struct{ ID gql.ID }) (*templateResolver, error) {
	localID, err := fromGlobalID("template", string(args.ID))
	if err != nil {
		return nil, err
	}
	result, err := res.dataSource.GetTemplate(localID)
	if err != nil {
		return nil, err
	}
	return &templateResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}

func (res *Resolver) Templates() *[]*templateResolver {
	var resolverList []*templateResolver
	templates, err := res.dataSource.GetTemplateList()
	if err != nil {
		panic(err)
	}
	for index := range templates {
		resolverList = append(
			resolverList,
			&templateResolver{
				dataSource: res.dataSource,
				data:       &templates[index],
			},
		)
	}
	return &resolverList
}

type updateTemplateArgs struct {
	ID    string
	Input struct {
		Name string
	}
}

func (res *Resolver) TemplateUpdate(args updateTemplateArgs) (*templateUpdateResultResolver, error) {
	localID, err := fromGlobalID("template", args.ID)
	if err != nil {
		return nil, err
	}
	if args.Input.Name == "" {
		return &templateUpdateResultResolver{
			dataSource: res.dataSource,
			data: templateUpdateResult{
				userErrors: &[]userError{
					userError{
						field:   "name",
						message: errNoEmpty("name").Error(),
					},
				},
				templateID: localID,
			},
		}, nil
	}
	input := core.TemplateInput{
		Name: args.Input.Name,
	}
	err = res.dataSource.UpdateTemplate(localID, input)
	if err != nil {
		return nil, err
	}
	return &templateUpdateResultResolver{
		dataSource: res.dataSource,
		data: templateUpdateResult{
			templateID: localID,
		},
	}, nil
}

type createTemplateFieldArgs struct {
	ID    gql.ID
	Input core.TemplateField
}

func (res *Resolver) CreateTemplateField(args createTemplateFieldArgs) (*templateUpdateResultResolver, error) {
	localID, err := fromGlobalID("template", string(args.ID))
	if err != nil {
		return nil, err
	}
	// TODO: Refactor this someday
	if args.Input.Name == "" {
		return &templateUpdateResultResolver{
			dataSource: res.dataSource,
			data: templateUpdateResult{
				userErrors: &[]userError{
					userError{
						field:   "name",
						message: errNoEmpty("name").Error(),
					},
				},
				templateID: localID,
			},
		}, nil
	}
	if args.Input.Type == "" {
		return &templateUpdateResultResolver{
			dataSource: res.dataSource,
			data: templateUpdateResult{
				userErrors: &[]userError{
					userError{
						field:   "type",
						message: errNoEmpty("type").Error(),
					},
				},
				templateID: localID,
			},
		}, nil
	}
	err = res.dataSource.AddTemplateField(localID, args.Input)
	if err != nil {
		return nil, err
	}
	return &templateUpdateResultResolver{
		dataSource: res.dataSource,
		data: templateUpdateResult{
			templateID: localID,
		},
	}, nil
}

type removeTemplateFieldArgs struct {
	ID    gql.ID
	Input struct {
		Name string
	}
}

func (res *Resolver) RemoveTemplateField(args removeTemplateFieldArgs) (*templateUpdateResultResolver, error) {
	localID, err := fromGlobalID("template", string(args.ID))
	if err != nil {
		return nil, err
	}
	err = res.dataSource.RemoveTemplateField(localID, args.Input.Name)
	if err != nil {
		return nil, err
	}
	return &templateUpdateResultResolver{
		dataSource: res.dataSource,
		data: templateUpdateResult{
			templateID: localID,
		},
	}, nil
}

type removeTemplateArgs struct {
	ID gql.ID
}

func (res *Resolver) RemoveTemplate(args removeTemplateArgs) (*templateRemoveResultResolver, error) {
	localID, err := fromGlobalID("template", string(args.ID))
	if err != nil {
		return nil, err
	}
	err = res.dataSource.RemoveTemplate(localID)
	if err != nil {
		return nil, err
	}
	return &templateRemoveResultResolver{
		dataSource: res.dataSource,
		data: templateRemoveResult{
			templateID: localID,
		},
	}, nil
}
