package art

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchArtObj *graphql.Object

// RootObject for art-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchArtObj != nil {
		return searchArtObj
	}

	// art
	searchArtObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "art",
		Description: "art-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(artSearchObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					dbStr := parse.BuildQuery(options, "art")
					entries := execute(dbStr, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchArtObj
}

var artSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "art_search",
	Fields: graphql.Fields{
		"art": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(art))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Art); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_art(): error")
			},
		},
	},
})
