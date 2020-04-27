package wallpaper

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

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
