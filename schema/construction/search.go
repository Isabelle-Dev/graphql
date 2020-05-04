package construction

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchConstructionObj *graphql.Object

// RootObject contains the main construction-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchConstructionObj != nil {
		return searchConstructionObj
	}

	// construction
	searchConstructionObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "construction",
		Description: "construction-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(constructionSearchObj),
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
					q := parse.BuildQuery(options, "construction", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchConstructionObj
}

var constructionSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "search_construction",
	Fields: graphql.Fields{
		"construction": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(construction)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]newhorizons.ConstructionEntry); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_construction(): error")
			},
		},
	},
})
