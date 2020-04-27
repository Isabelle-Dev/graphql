package photos

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

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
