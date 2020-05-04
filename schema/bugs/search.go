package bugs

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchBugObj *graphql.Object

// RootObject contains the main bug-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchBugObj != nil {
		return searchBugObj
	}

	// bug
	searchBugObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "bugs",
		Description: "bug-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(bugSearchObj),
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
					q := parse.BuildQuery(options, "bug", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchBugObj
}

var bugSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "bug_search",
	Fields: graphql.Fields{
		"bugs": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(bug)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Bug); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_bug(): error")
			},
		},
	},
})
