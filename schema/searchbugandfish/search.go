package searchbugandfish

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var bugAndFishSearchObject *graphql.Object

// RootObject that contains search related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if bugAndFishSearchObject != nil {
		return bugAndFishSearchObject
	}

	// Search
	bugAndFishSearchObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "bug_and_fish",
		Description: "Bug and Fish-related Sources",
		Fields: graphql.Fields{
			"search_bug_north": &graphql.Field{
				Name: "Search a bug from northern hemisphere",
				Type: graphql.NewNonNull(searchBugResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)
					entry := findBugByName(item, "north_bug", db)
					if entry == nil {
						return nil, fmt.Errorf("search(): %v not found", item)
					}
					return entry, nil
				},
			},
			"search_fish_north": &graphql.Field{
				Name: "Search a fish from northern hemisphere",
				Type: graphql.NewNonNull(searchFishResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)

					entry := findFishByName(item, "north_fish", db)
					if entry == nil {
						return nil, fmt.Errorf("search(): %v not found", item)
					}
					return entry, nil
				},
			},
			"search_bug_south": &graphql.Field{
				Name: "Search a bug from southern hemisphere",
				Type: graphql.NewNonNull(searchBugResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)

					entry := findBugByName(item, "south_bug", db)
					if entry == nil {
						return nil, fmt.Errorf("search(): %v not found", item)
					}
					return entry, nil
				},
			},
			"search_fish_south": &graphql.Field{
				Name: "Search a fish from southern hemisphere",
				Type: graphql.NewNonNull(searchFishResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)

					entry := findFishByName(item, "south_fish", db)
					if entry == nil {
						return nil, fmt.Errorf("search(): %v not found", item)
					}
					return entry, nil
				},
			},

			// "search_all_by_north_hemisphere": &graphql.Field{
			// 	Name: "Search all bug and fish in the north",
			// 	Type: graphql.NewNonNull(searchCombinedObj),
			// 	Args: graphql.FieldConfigArgument{
			// 		"hemi": &graphql.ArgumentConfig{
			// 			Type: graphql.String,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			// 		hemi := p.Args["hemi"].(string)
			// 		fmt.Println(hemi)

			// 	},
			// },
		},
	})
	return bugAndFishSearchObject
}
