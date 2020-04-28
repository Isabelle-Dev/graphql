package poster

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var posterSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "poster_search",
	Fields: graphql.Fields{
		"posters": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(poster))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Poster); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_posters(): error")
			},
		},
	},
})
