package clothes

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/Isabelle-Dev/graphql/parse"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var searchClothesObj *graphql.Object

// RootObject contains the main clothes-related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if searchClothesObj != nil {
		return searchClothesObj
	}

	// clothes
	searchClothesObj = graphql.NewObject(graphql.ObjectConfig{
		Name:        "clothes",
		Description: "clothes-related query sources",
		Fields: graphql.Fields{
			"query": &graphql.Field{
				Name: "query fields",
				Type: graphql.NewNonNull(clothesSearchObj),
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
					q := parse.BuildQuery(options, "clothes", limit, glob)
					entries := execute(q, db)
					if entries == nil {
						return nil, fmt.Errorf("query(): no entries found")
					}
					return entries, nil
				},
			},
		},
	})

	return searchClothesObj
}

var clothesSearchObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "search_clothes",
	Fields: graphql.Fields{
		"clothes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(clothes)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.([]*newhorizons.Clothes); ok {
					return val, nil
				}
				return nil, fmt.Errorf("search_clothes(): error")
			},
		},
	},
})
