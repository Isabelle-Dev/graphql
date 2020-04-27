package floor

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var floorSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "floor_search",
	Fields: graphql.Fields{
		"floors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(floor))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Floor); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_floor(): error")
			},
		},
	},
})
