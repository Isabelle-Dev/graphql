package photos

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchPhotoObj *graphql.Object

// RootObject for photo-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchPhotoObj != nil {
		return searchPhotoObj
	}

	// photo
	searchPhotoObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "photo",
		Description: "Villager photo-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(photoSearchObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					dbStr := parse.BuildQuery(options, "photos")
					entries := execute(dbStr, db)
					if entries == nil || len(entries) == 0 {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchPhotoObj
}

var photoSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "photo_search",
	Fields: graphql.Fields{
		"photos": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(photo))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Photo); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_photos(): error")
			},
		},
	},
})
