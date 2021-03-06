package item

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchItemObj *graphql.Object

// RootObject contains the main item-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchItemObj != nil {
		return searchItemObj
	}

	// item
	searchItemObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "item",
		Description: "item-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(itemSearchObj),
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
						DefaultValue: 1000,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					glob := p.Args["glob"].(string)
					limit := p.Args["limit"].(int)
					options := parse.QueryParse(query)
					q := parse.BuildQuery(options, "item", limit, glob)
					// Debugging purposes
					// fmt.Println(q)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchItemObj
}

var itemSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "item_search",
	Fields: graphql.Fields{
		"items": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(item))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Item); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_item(): Error")
			},
		},
	},
})
