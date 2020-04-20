package searchbugandfish

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchCombinedObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "AllSearchResultsByHemi",
	Fields: graphql.Fields{
		"bugs": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(combinedBugsObject))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Combined); ok {
					return val.Bugs, nil
				}
				return nil, fmt.Errorf("search(): Not a recognized Combined object")
			},
		},
		"fishes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(combinedFishesObject))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Combined); ok {
					return val.Fishes, nil
				}
				return nil, fmt.Errorf("search(): Not a recognized Combined object")
			},
		},
	},
})

var combinedBugsObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AllBugsInHemi",
	Description: "All bugs available currently in hemisphere",
	Fields: graphql.Fields{
		"bug": &graphql.Field{
			Type: graphql.NewNonNull(searchBugResObj),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.BugEntry); ok {
					return &val, nil
				}
				return nil, fmt.Errorf("search(): couldn't resolve bug entry")
			},
		},
	},
})

var combinedFishesObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AllFishesInHemi",
	Description: "All fishes available currently in hemisphere",
	Fields: graphql.Fields{
		"fish": &graphql.Field{
			Type: graphql.NewNonNull(searchFishResObj),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.FishEntry); ok {
					return &val, nil
				}
				return nil, fmt.Errorf("search(): couldn't resolve fish entry")
			},
		},
	},
})
