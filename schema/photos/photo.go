package photos

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var photo = graphql.NewObject(graphql.ObjectConfig{
	Name: "photoObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Variants": &graphql.Field{
			Type: graphql.NewList(variantObj),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.Variants, nil
				}
				return nil, nil
			},
		},
		"BodyTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.BodyTitle, nil
				}
				return nil, nil
			},
		},
		"PatternTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.PatternTitle, nil
				}
				return nil, nil
			},
		},
		"DIY": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.DIY, nil
				}
				return nil, nil
			},
		},
		"Customize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.Customize, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"KitCost": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.KitCost, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Photo); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
	},
})

var variantObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "variantObj",
	Fields: graphql.Fields{
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Variant); ok {
					return val.ImageURL, nil
				}
				return nil, nil
			},
		},
		"Pattern": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Variant); ok {
					return val.Pattern, nil
				}
				return nil, nil
			},
		},
		"Colors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Variant); ok {
					return val.Colors, nil
				}
				return nil, nil
			},
		},
	},
})
