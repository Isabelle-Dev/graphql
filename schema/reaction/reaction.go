package reaction

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var reaction = graphql.NewObject(graphql.ObjectConfig{
	Name: "reactionObj",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ReactionEntry); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ReactionEntry); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.ReactionEntry); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
	},
})
