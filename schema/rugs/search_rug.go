package rug

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchRugObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "rug_search",
	Fields: graphql.Fields{
		"rugs": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(rug))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Rug); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_rug(): error")
			},
		},
	},
})
