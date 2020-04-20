package searchitem

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchItemObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "search_by_item",
	Fields: graphql.Fields{
		"items": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(itemObj))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*[]newhorizons.Item); ok {
					return *val, nil
				}
				return nil, fmt.Errorf("search_by_item(): Error - not a recognized type")
			},
		},
	},
})

var itemObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndividualItem",
	Fields: graphql.Fields{
		"item": &graphql.Field{
			Type: graphql.NewNonNull(item),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val, nil
				}
				return nil, fmt.Errorf("IndividualItem(): error")
			},
		},
	},
})
