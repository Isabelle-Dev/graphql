package bugandfish

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var searchFishResObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "SearchFishResult",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.Image, nil
				}
				return nil, nil
			},
		},
		"House": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.House, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.Sell, nil
				}
				return nil, nil
			},
		},
		"Where": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.Where, nil
				}
				return nil, nil
			},
		},
		"Shadow": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.Shadow, nil
				}
				return nil, nil
			},
		},
		"Rarity": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.Rarity, nil
				}
				return nil, nil
			},
		},
		"RainSnowUp": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.RainSnowUp, nil
				}
				return nil, nil
			},
		},
		"StartTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.StartTime, nil
				}
				return nil, nil
			},
		},
		"EndTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.EndTime, nil
				}
				return nil, nil
			},
		},
		"NorthernHemi": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.NorthernHemi.Months, nil
				}
				return nil, nil
			},
		},
		"SouthernHemi": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.Fish); ok {
					return entry.SouthernHemi.Months, nil
				}
				return nil, nil
			},
		},
	},
})
