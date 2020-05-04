package fish

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var fish = graphql.NewObject(graphql.ObjectConfig{
	Name: "fishObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Location": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Location, nil
				}
				return nil, nil
			},
		},
		"Shadow": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Shadow, nil
				}
				return nil, nil
			},
		},
		"Rarity": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Rarity, nil
				}
				return nil, nil
			},
		},
		"CatchUp": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.CatchUp, nil
				}
				return nil, nil
			},
		},
		"StartTime": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.StartTime, nil
				}
				return nil, nil
			},
		},
		"EndTime": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.EndTime, nil
				}
				return nil, nil
			},
		},
		"Colors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Fish); ok {
					return val.Color, nil
				}
				return nil, nil
			},
		},
	},
})
