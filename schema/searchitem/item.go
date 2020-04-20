package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var item = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemResult",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Variation": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Variation, nil
				}
				return nil, nil
			},
		},
		"BodyTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.BodyTitle, nil
				}
				return nil, nil
			},
		},
		"Pattern": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Pattern, nil
				}
				return nil, nil
			},
		},
		"PatternTitle": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.PatternTitle, nil
				}
				return nil, nil
			},
		},
		"DIY": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.DIY, nil
				}
				return nil, nil
			},
		},
		"BodyCustomize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.BodyCustomize, nil
				}
				return nil, nil
			},
		},
		"PatternCustomize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.PatternCustomize, nil
				}
				return nil, nil
			},
		},
		"KitCost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.KitCost, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Color1": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Color1, nil
				}
				return nil, nil
			},
		},
		"Color2": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Color2, nil
				}
				return nil, nil
			},
		},
		"Size": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Size, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"SourceNotes": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.SourceNotes, nil
				}
				return nil, nil
			},
		},
		"Version": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Version, nil
				}
				return nil, nil
			},
		},
		"HHAConcept1": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.HHAConcept1, nil
				}
				return nil, nil
			},
		},
		"HHAConcept2": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.HHAConcept2, nil
				}
				return nil, nil
			},
		},
		"HHASeries": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.HHASeries, nil
				}
				return nil, nil
			},
		},
		"HHASet": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.HHASet, nil
				}
				return nil, nil
			},
		},
		"Interact": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Interact, nil
				}
				return nil, nil
			},
		},
		"Tag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Tag, nil
				}
				return nil, nil
			},
		},
		"SpeakerType": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.SpeakerType, nil
				}
				return nil, nil
			},
		},
		"LightingType": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.LightingType, nil
				}
				return nil, nil
			},
		},
		"Catalog": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Catalog, nil
				}
				return nil, nil
			},
		},
		"Filename": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.Filename, nil
				}
				return nil, nil
			},
		},
		"VariantID": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.VariantID, nil
				}
				return nil, nil
			},
		},
		"InternalID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Item); ok {
					return val.InternalID, nil
				}
				return nil, nil
			},
		},
	},
})
