package search

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchObject *graphql.Object

func RootObject(db *gorm.DB) *graphql.Object {
	if searchObject != nil {
		return searchObject
	}

	// Search
	searchObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Search",
		Description: "Search-related Sources",
		Fields: graphql.Fields{
			"search_bug_or_fish": &graphql.Field{
				Name: "Search a bug or fish",
				Type: graphql.NewNonNull(searchBugOrFishResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"tablename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)
					tableName := p.Args["tablename"].(string)

					entry := findByName(item, tableName, db)
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
