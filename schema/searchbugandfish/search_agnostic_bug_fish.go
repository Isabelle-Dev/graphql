package searchbugandfish

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var agnosticBugObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "SearchAgBugResult",
	Fields: graphql.Fields{
		"Num": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Num, nil
				}
				return nil, nil
			},
		},
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Image, nil
				}
				return nil, nil
			},
		},
		"House": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.House, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Sell, nil
				}
				return nil, nil
			},
		},
		"Where": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Where, nil
				}
				return nil, nil
			},
		},
		"Weather": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Weather, nil
				}
				return nil, nil
			},
		},
		"Rarity": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Rarity, nil
				}
				return nil, nil
			},
		},
		"StartTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.StartTime, nil
				}
				return nil, nil
			},
		},
		"EndTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.EndTime, nil
				}
				return nil, nil
			},
		},
		"Color1": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Color1, nil
				}
				return nil, nil
			},
		},
		"Color2": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.Color2, nil
				}
				return nil, nil
			},
		},
		"ItemFilename": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.ItemFilename, nil
				}
				return nil, nil
			},
		},
		"InternalID": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.BugTimeAgnostic); ok {
					return entry.InternalID, nil
				}
				return nil, nil
			},
		},
	},
})

var agnosticFishObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "SearchAgFishResult",
	Fields: graphql.Fields{
		"Num": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Num, nil
				}
				return nil, nil
			},
		},
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Image, nil
				}
				return nil, nil
			},
		},
		"House": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.House, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Sell, nil
				}
				return nil, nil
			},
		},
		"Where": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Where, nil
				}
				return nil, nil
			},
		},
		"Shadow": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Shadow, nil
				}
				return nil, nil
			},
		},
		"Rarity": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Rarity, nil
				}
				return nil, nil
			},
		},
		"RainSnowUp": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.RainSnowUp, nil
				}
				return nil, nil
			},
		},
		"StartTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.StartTime, nil
				}
				return nil, nil
			},
		},
		"EndTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.EndTime, nil
				}
				return nil, nil
			},
		},
		"Color1": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Color1, nil
				}
				return nil, nil
			},
		},
		"Color2": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Color2, nil
				}
				return nil, nil
			},
		},
		"Size": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.Size, nil
				}
				return nil, nil
			},
		},
		"LightingType": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.LightingType, nil
				}
				return nil, nil
			},
		},
		"ItemFilename": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.ItemFilename, nil
				}
				return nil, nil
			},
		},
		"InternalID": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*newhorizons.FishTimeAgnostic); ok {
					return entry.InternalID, nil
				}
				return nil, nil
			},
		},
	},
})
