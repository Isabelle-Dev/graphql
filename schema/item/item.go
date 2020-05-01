package item

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var item = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemResult",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Variants": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(variantObj)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Variants, nil
				}
				return nil, nil
			},
		},
		"BodyTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.BodyTitle, nil
				}
				return nil, nil
			},
		},
		"PatternTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.PatternTitle, nil
				}
				return nil, nil
			},
		},
		"DIY": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.DIY, nil
				}
				return nil, nil
			},
		},
		"BodyCustomize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.BodyCustomize, nil
				}
				return nil, nil
			},
		},
		"PatternCustomize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.PatternCustomize, nil
				}
				return nil, nil
			},
		},
		"KitCost": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.KitCost, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"Concepts": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.HHAConcepts, nil
				}
				return nil, nil
			},
		},
		"HHASeries": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.HHASeries, nil
				}
				return nil, nil
			},
		},
		"HHASet": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.HHASet, nil
				}
				return nil, nil
			},
		},
		"Interact": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Interact, nil
				}
				return nil, nil
			},
		},
		"Tag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Item); ok {
					return val.Tag, nil
				}
				return nil, nil
			},
		},
	},
})

var variantObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Variant",
	Fields: graphql.Fields{
		"Img": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Variant); ok {
					return val.Image, nil
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
