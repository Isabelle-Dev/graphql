package music

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var music = graphql.NewObject(graphql.ObjectConfig{
	Name: "musicObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"SourceNotes": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.SourceNotes, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"Colors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Music); ok {
					return val.Color, nil
				}
				return nil, nil
			},
		},
	},
})
