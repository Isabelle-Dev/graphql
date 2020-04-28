package art

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var art = graphql.NewObject(graphql.ObjectConfig{
	Name: "artObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Type": &graphql.Field{
			Type: graphql.NewList(artType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Type, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"Concept": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Concept, nil
				}
				return nil, nil
			},
		},
		"Tag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Tag, nil
				}
				return nil, nil
			},
		},
		"Category": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Art); ok {
					return val.Category, nil
				}
				return nil, nil
			},
		},
	},
})

var artType = graphql.NewObject(graphql.ObjectConfig{
	Name: "artType",
	Fields: graphql.Fields{
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ArtType); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ArtType); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Genuine": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ArtType); ok {
					return val.Genuine, nil
				}
				return nil, nil
			},
		},
	},
})
