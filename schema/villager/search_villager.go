package villager

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var villagerSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "villager_search",
	Fields: graphql.Fields{
		"villagers": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(villager))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Villager); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_villager(): error")
			},
		},
	},
})
