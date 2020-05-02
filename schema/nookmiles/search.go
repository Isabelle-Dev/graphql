package nookmiles

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchNookMiles *graphql.Object

// RootObject contains the main nook mile-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchNookMiles != nil {
		return searchNookMiles
	}

	// nook miles
	searchNookMiles = graphql.NewObject(graphql.ObjectConfig{
		Name:        "nookmiles",
		Description: "nook miles-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(nookmileSearch),
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
					q := parse.BuildQuery(options, "floor", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchNookMiles
}

var nookmileSearch = graphql.NewObject(graphql.ObjectConfig{
	Name: "nookmile_search",
	Fields: graphql.Fields{
		"nook_mile": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(nookmile))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]newhorizons.NookMiles); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_nookmile(): error")
			},
		},
	},
})
