package searchclothes

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var clothesSearchObject *graphql.Object

// RootObject contains the main clothes-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if clothesSearchObject != nil {
		return clothesSearchObject
	}

	// clothes
	clothesSearchObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "clothes",
		Description: "clothes-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(searchClothesObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					q := parse.BuildQuery(options, "clothes")
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return clothesSearchObject
}
