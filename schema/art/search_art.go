package art

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var artSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "art_search",
	Fields: graphql.Fields{
		"art": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(art))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Art); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_art(): error")
			},
		},
	},
})
