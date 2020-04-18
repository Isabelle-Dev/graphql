package search

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchObject *graphql.Object

// RootObject that contains search related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchObject != nil {
		return searchObject
	}

	// Search
	searchObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Search",
		Description: "Search-related Sources",
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
		},
	})
	return searchObject
}
