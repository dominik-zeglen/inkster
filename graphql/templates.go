package graphql

import (
	"github.com/dominik-zeglen/ecoknow/core"
	// "github.com/globalsign/mgo/bson"
	// gql "github.com/graph-gophers/graphql-go"
)

// Type resolvers
type templateResolver struct {
	dataSource core.Adapter
	data       *core.Template
}
type templateFieldResolver struct {
	dataSource core.Adapter
	data       *core.TemplateField
}

func (res *templateFieldResolver) Name() string {
	return res.data.Name
}

func (res *templateFieldResolver) Type() string {
	return res.data.Type
}

func (res *templateResolver) ID() string {
	globalID := toGlobalID("template", res.data.ID)
	return globalID
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

func (res *Resolver) Template(args struct{ ID string }) (*templateResolver, error) {
	localID, err := fromGlobalID("template", args.ID)
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

// func (res *Resolver) TemplateUpdate(args updateTemplateArgs) (*templateOperationResolver, error) {
// }
