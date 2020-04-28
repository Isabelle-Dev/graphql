package poster

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchPosterObj *graphql.Object

// RootObject for poster-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchPosterObj != nil {
		return searchPosterObj
	}

	// poster
	searchPosterObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "posters",
		Description: "Villager poster-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(posterSearchObj),
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
					dbStr := parse.BuildQuery(options, "poster", limit, glob)
					entries := execute(dbStr, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchPosterObj
}

var posterSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "poster_search",
	Fields: graphql.Fields{
		"posters": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(poster))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Poster); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_posters(): error")
			},
		},
	},
})
