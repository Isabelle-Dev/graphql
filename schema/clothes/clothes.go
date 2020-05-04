package clothes

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var clothes = graphql.NewObject(graphql.ObjectConfig{
	Name: "clothesObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Variants": &graphql.Field{
			Type: graphql.NewList(clothesVariant),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.Variants, nil
				}
				return nil, nil
			},
		},
		"DIY": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.DIY, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"SourceNotes": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.SourceNotes, nil
				}
				return nil, nil
			},
		},
		"Themes": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Clothes); ok {
					return val.Themes, nil
				}
				return nil, nil
			},
		},
	},
})

var clothesVariant = graphql.NewObject(graphql.ObjectConfig{
	Name: "clothesvariants",
	Fields: graphql.Fields{
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ClothesVariant); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Variation": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ClothesVariant); ok {
					return val.Variation, nil
				}
				return nil, nil
			},
		},
		"Colors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ClothesVariant); ok {
					return val.Color, nil
				}
				return nil, nil
			},
		},
	},
})
