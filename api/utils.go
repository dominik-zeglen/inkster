package api

import "github.com/dominik-zeglen/inkster/core"

func createInputErrorResolvers(errors []core.ValidationError) []inputErrorResolver {
	var resolverList []inputErrorResolver
	for i := range errors {
		resolverList = append(
			resolverList,
			inputErrorResolver{
				err: errors[i],
			},
		)
	}

	return resolverList
}
