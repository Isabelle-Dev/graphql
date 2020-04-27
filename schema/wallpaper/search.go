package wallpaper

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var wallpaperSearchObj *graphql.Object

// RootObject (wallpaper) contains wallpaper-related search queries
func RootObject(db *gorm.DB) *graphql.Object {
	if wallpaperSearchObj != nil {
		return wallpaperSearchObj
	}

	wallpaperSearchObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "wallpapers",
		Description: "Wallpaper-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(searchWallpaperObj),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					q := parse.BuildQuery(options, "wallpaper")
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return wallpaperSearchObj
}
