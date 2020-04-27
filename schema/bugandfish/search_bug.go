package bugandfish

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchBugResObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "SearchBugResult",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.Image, nil
				}
				return nil, nil
			},
		},
		"House": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.House, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.Sell, nil
				}
				return nil, nil
			},
		},
		"Where": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.Where, nil
				}
				return nil, nil
			},
		},
		"Weather": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.Weather, nil
				}
				return nil, nil
			},
		},
		"Rarity": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.Rarity, nil
				}
				return nil, nil
			},
		},
		"StartTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.StartTime, nil
				}
				return nil, nil
			},
		},
		"EndTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.EndTime, nil
				}
				return nil, nil
			},
		},
		"NorthernHemi": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.NorthernHemi.Months, nil
				}
				return nil, nil
			},
		},
		"SouthernHemi": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Bug); ok {
					return entry.SouthernHemi.Months, nil
				}
				return nil, nil
			},
		},
	},
})
