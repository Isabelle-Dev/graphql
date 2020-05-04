package tool

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var tool = graphql.NewObject(graphql.ObjectConfig{
	Name: "toolObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Variants": &graphql.Field{
			Type: graphql.NewList(toolVariant),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Variant, nil
				}
				return nil, nil
			},
		},
		"BodyTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.BodyTitle, nil
				}
				return nil, nil
			},
		},
		"DIY": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.DIY, nil
				}
				return nil, nil
			},
		},
		"Customize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Customize, nil
				}
				return nil, nil
			},
		},
		"KitCost": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.KitCost, nil
				}
				return nil, nil
			},
		},
		"Uses": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Uses, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Set": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Set, nil
				}
				return nil, nil
			},
		},
		"Sources": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"SourceNotes": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Tool); ok {
					return val.SourceNotes, nil
				}
				return nil, nil
			},
		},
	},
})

var toolVariant = graphql.NewObject(graphql.ObjectConfig{
	Name: "tool_variant",
	Fields: graphql.Fields{
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ToolVariant); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Variation": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ToolVariant); ok {
					return val.Variation, nil
				}
				return nil, nil
			},
		},
	},
})
