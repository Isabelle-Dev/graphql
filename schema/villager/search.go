package villager

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/Isabelle-Dev/isabelle-graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchVillagerObj *graphql.Object

// RootObject contains the main villager-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchVillagerObj != nil {
		return searchVillagerObj
	}

	// villager
	searchVillagerObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "villager",
		Description: "villager-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(villagerSearchObj),
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
					q := parse.BuildQuery(options, "villagers", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchVillagerObj
}

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
