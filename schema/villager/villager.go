package villager

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var villager = graphql.NewObject(graphql.ObjectConfig{
	Name: "villagerObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Species": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Species, nil
				}
				return nil, nil
			},
		},
		"Gender": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Gender, nil
				}
				return nil, nil
			},
		},
		"Personality": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Personality, nil
				}
				return nil, nil
			},
		},
		"Birthday": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Birthday, nil
				}
				return nil, nil
			},
		},
		"Catchphrase": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Catchphrase, nil
				}
				return nil, nil
			},
		},
		"Style": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Style, nil
				}
				return nil, nil
			},
		},
		"Color": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Villager); ok {
					return val.Color, nil
				}
				return nil, nil
			},
		},
	},
})
