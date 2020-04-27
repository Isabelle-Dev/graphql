package villager

import (
	"fmt"

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
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					query := p.Args["query"].(string)
					options := parse.QueryParse(query)
					q := parse.BuildQuery(options, "villagers")
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
