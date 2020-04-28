package rug

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var rugSearchObj *graphql.Object

// RootObject for rug-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if rugSearchObj != nil {
		return rugSearchObj
	}

	// rug
	rugSearchObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "rug",
		Description: "Rug-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(searchRugObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					dbStr := parse.BuildQuery(options, "rugs")
					entries := execute(dbStr, db)
					if entries == nil || len(entries) == 0 {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return rugSearchObj
}

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
