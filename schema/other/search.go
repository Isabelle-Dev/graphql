package other

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchOtherObj *graphql.Object

// RootObject contains the main 'other'-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchOtherObj != nil {
		return searchOtherObj
	}

	// other
	searchOtherObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "other",
		Description: "other-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(otherSearchObj),
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
					q := parse.BuildQuery(options, "other", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchOtherObj
}

var otherSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "other_search",
	Fields: graphql.Fields{
		"others": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(other))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]newhorizons.Other); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_other(): error")
			},
		},
	},
})
