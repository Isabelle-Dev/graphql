package searchbugandfish

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchAgnosticObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "AllSearchByAgnosticTime",
	Fields: graphql.Fields{
		"bugs": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(agCombinedBugsObj))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.CombinedAgnostic); ok {
					return val.Bugs, nil
				}
				return nil, fmt.Errorf("Error")
			},
		},
		"fishes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(agCombinedFishObj))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.CombinedAgnostic); ok {
					return val.Fishes, nil
				}
				return nil, fmt.Errorf("error")
			},
		},
	},
})

var agCombinedBugsObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndividualAgnosticBug",
	Fields: graphql.Fields{
		"bug": &graphql.Field{
			Type: graphql.NewNonNull(agnosticBugObj),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.BugTimeAgnostic); ok {
					return &val, nil
				}
				return nil, fmt.Errorf("search(): couldn't resolve bug entry")
			},
		},
	},
})

var agCombinedFishObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndividualAgnosticFish",
	Fields: graphql.Fields{
		"fish": &graphql.Field{
			Type: graphql.NewNonNull(agnosticFishObj),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.FishTimeAgnostic); ok {
					return &val, nil
				}
				return nil, fmt.Errorf("search(): couldn't resolve bug entry")
			},
		},
	},
})
