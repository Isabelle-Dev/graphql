package music

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchMusicObj *graphql.Object

// RootObject for music-related type queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchMusicObj != nil {
		return searchMusicObj
	}

	// music
	searchMusicObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "music",
		Description: "music-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(musicSearchObj),
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
					dbStr := parse.BuildQuery(options, "music", limit, glob)
					entries := execute(dbStr, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchMusicObj
}

var musicSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "music_search",
	Fields: graphql.Fields{
		"music": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(music))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Music); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_rug(): error")
			},
		},
	},
})
