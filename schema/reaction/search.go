package reaction

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchReactionObj *graphql.Object

// RootObject contains the main reaction-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchReactionObj != nil {
		return searchReactionObj
	}

	// reaction
	searchReactionObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "reaction",
		Description: "reaction-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(reactionSearchObj),
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
					q := parse.BuildQuery(options, "reaction", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchReactionObj
}

var reactionSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "reaction_search",
	Fields: graphql.Fields{
		"reactions": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(reaction)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]newhorizons.ReactionEntry); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_reaction(): error")
			},
		},
	},
})
