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
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CombinedBugsObject))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Combined); ok {
					return val.Bugs, nil
				}
				return nil, fmt.Errorf("search(): Not a recognized combine struct")
			},
		},
		"fishes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CombinedFishesObject))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Combined); ok {
					return val.Fishes, nil
				}
				return nil, fmt.Errorf("search(): Not a recognized combine struct")
			},
		},
	},
})

// CombinedBugsObject takes a bug entry query, resolves it, and returns the value
var CombinedBugsObject = graphql.NewObject(graphql.ObjectConfig{
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

// CombinedFishesObject takes a fish entry query, resolves it, and returns the value
var CombinedFishesObject = graphql.NewObject(graphql.ObjectConfig{
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
