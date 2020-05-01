package fencing

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchFenceObject *graphql.Object

// RootObject contains the main fencing-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchFenceObject != nil {
		return searchFenceObject
	}

	// fence
	searchFenceObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "fencing",
		Description: "fence-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(fenceSearchObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"glob": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "f",
					},
					"limit": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 500,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					glob := p.Args["glob"].(string)
					limit := p.Args["limit"].(int)
					options := parse.QueryParse(query)
					q := parse.BuildQuery(options, "floor", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchFenceObject
}

var fenceSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "fence_search",
	Fields: graphql.Fields{
		"fences": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(fence)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]newhorizons.FenceEntry); ok {
					return val, nil
				}
				return nil, fmt.Errorf("fence_search(): error")
			},
		},
	},
})
