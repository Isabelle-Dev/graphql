package umbrella

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchUmbrellaObj *graphql.Object

// RootObject contains the main umbrella-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchUmbrellaObj != nil {
		return searchUmbrellaObj
	}

	// umbrella
	searchUmbrellaObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "umbrellas",
		Description: "umbrella-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(umbrellaSearchObj),
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
					q := parse.BuildQuery(options, "umbrella", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchUmbrellaObj
}

var umbrellaSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "umbrella_search",
	Fields: graphql.Fields{
		"umbrellas": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(umbrella))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Umbrella); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_umbrella(): error")
			},
		},
	},
})
