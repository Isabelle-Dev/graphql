package nookmiles

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var nookmile = graphql.NewObject(graphql.ObjectConfig{
	Name: "nook_mileObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.NookMiles); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.NookMiles); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"NookMiles": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.NookMiles); ok {
					return val.NookMiles, nil
				}
				return nil, nil
			},
		},
		"Category": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.NookMiles); ok {
					return val.Category, nil
				}
				return nil, nil
			},
		},
	},
})
