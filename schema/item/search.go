package item

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var itemSearchObject *graphql.Object

// RootObject contains the main item-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if itemSearchObject != nil {
		return itemSearchObject
	}

	// item
	itemSearchObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "item",
		Description: "item-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(searchItemObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					q := parse.BuildQuery(options, "item")
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

	return itemSearchObject
}

var searchItemObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "search_by_item",
	Fields: graphql.Fields{
		"items": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(item))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Item); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_by_item(): Error - not a recognized type")
			},
		},
	},
})
