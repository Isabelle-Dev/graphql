package tool

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchToolObj *graphql.Object

// RootObject for photo-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchToolObj != nil {
		return searchToolObj
	}

	// tool
	searchToolObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "tool",
		Description: "Tool-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(toolSearchObj),
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
					dbStr := parse.BuildQuery(options, "tools", limit, glob)
					entries := execute(dbStr, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchToolObj
}

var toolSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "tool_search",
	Fields: graphql.Fields{
		"tools": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(tool))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Tool); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_tools(): error")
			},
		},
	},
})
