package wallpaper

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
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
					q := parse.BuildQuery(options, "wallpaper", limit, glob)
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

var searchWallpaperObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "search_wallpaper",
	Fields: graphql.Fields{
		"wallpapers": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(wallpaper))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Wallpaper); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_wallpaper(): Error")
			},
		},
	},
})
