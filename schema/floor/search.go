package floor

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchFloorObj *graphql.Object

// RootObject contains the main floor-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchFloorObj != nil {
		return searchFloorObj
	}

	// flooring
	searchFloorObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "flooring",
		Description: "flooring-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(floorSearchObj),
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

	return searchFloorObj
}

var floorSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "floor_search",
	Fields: graphql.Fields{
		"floors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(floor))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Floor); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_floor(): error")
			},
		},
	},
})
