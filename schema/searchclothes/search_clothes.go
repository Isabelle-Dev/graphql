package searchclothes

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchClothesObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "search_clothes",
	Fields: graphql.Fields{
		"clothes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(clothes)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Clothes); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_clothes(): error - not a recognized type")
			},
		},
	},
})
