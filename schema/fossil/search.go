package fossil

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchFossilObj *graphql.Object

// RootObject contains the main fossil-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchFossilObj != nil {
		return searchFossilObj
	}

	// fossil
	searchFossilObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "fossil",
		Description: "fossil-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(fossilSearchObj),
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
					q := parse.BuildQuery(options, "fossil", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchFossilObj
}

var fossilSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "fossil_search",
	Fields: graphql.Fields{
		"fossils": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(fossil)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Fossil); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_fossil(): error")
			},
		},
	},
})
