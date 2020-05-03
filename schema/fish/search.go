package fish

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchFishObj *graphql.Object

// RootObject for fish-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchFishObj != nil {
		return searchFishObj
	}

	// fish
	searchFishObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "fishes",
		Description: "fish-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(fishSearchObj),
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
					dbStr := parse.BuildQuery(options, "fish", limit, glob)
					entries := execute(dbStr, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchFishObj
}

var fishSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "fish_search",
	Fields: graphql.Fields{
		"fishes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(fish))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Fish); ok {
					return val, nil
				}
				return nil, fmt.Errorf("fish_search(): error")
			},
		},
	},
})
