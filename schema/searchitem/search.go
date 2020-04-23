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
				Type: graphql.NewNonNull(item),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)
					entry := findByName(item, "item", db)
					if entry == nil {
						return nil, fmt.Errorf("search_item(): item not found")
					}
					return entry, nil
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
					entries := findByColor(color, "item", db)
					if entries == nil {
						return nil, fmt.Errorf("search_color(): items not found")
					}
					return entries, nil
				},
			},

			"search_by_concept": &graphql.Field{
				Name: "search_by_concpet",
				Type: graphql.NewNonNull(searchItemObj),
				Args: graphql.FieldConfigArgument{
					"concept": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					concept := p.Args["concept"].(string)
					entries := findByConcept(concept, "item", db)
					if entries == nil {
						return nil, fmt.Errorf("search_concept(): items not found")
					}
					return entries, nil
				},
			},

			"only_door_hanging": &graphql.Field{
				Name: "door hanging objects",
				Type: graphql.NewNonNull(searchItemObj),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					entries := findAllDoorHanging("item", db)
					if entries == nil {
						return nil, fmt.Errorf("search_hanging(): items not found")
					}
					return entries, nil
				},
			},
		},
	})

	return itemSearchObject
}
