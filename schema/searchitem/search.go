package searchitem

import (
	"fmt"

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
			"search_item": &graphql.Field{
				Name: "search_by_item_name",
				Type: graphql.NewNonNull(searchItemObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)
					entries := findByName(item, "housewares", db)
					if entries == nil {
						return nil, fmt.Errorf("search_item(): error retrieving entries")
					}
					return entries, nil
				},
			},

			"search_by_color_variant": &graphql.Field{
				Name: "search_by_color_variant",
				Type: graphql.NewNonNull(searchItemObj),
				Args: graphql.FieldConfigArgument{
					"color": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					color := p.Args["color"].(string)
					entries := findByColor(color, "housewares", db)
					if entries == nil {
						return nil, fmt.Errorf("search_color(): error retrieving entries")
					}
					return entries, nil
				},
			},
		},
	})

	return itemSearchObject
}
