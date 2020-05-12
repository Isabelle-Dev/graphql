package bugs

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var bug = graphql.NewObject(graphql.ObjectConfig{
	Name: "bugObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Location": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.Location, nil
				}
				return nil, nil
			},
		},
		"Weather": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.Weather, nil
				}
				return nil, nil
			},
		},
		"TotalCatches": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.TotalCatches, nil
				}
				return nil, nil
			},
		},
		"MonthTime": &graphql.Field{
			Type: graphql.NewList(month),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.MonthTime, nil
				}
				return nil, nil
			},
		},
		"Colors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Bug); ok {
					return val.Color, nil
				}
				return nil, nil
			},
		},
	},
})

var month = graphql.NewObject(graphql.ObjectConfig{
	Name: "month_bug",
	Fields: graphql.Fields{
		"Month": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Month); ok {
					return val.Month, nil
				}
				return nil, nil
			},
		},
		"Time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Month); ok {
					return val.Time, nil
				}
				return nil, nil
			},
		},
	},
})
