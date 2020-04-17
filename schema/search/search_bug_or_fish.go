package search

import "github.com/graphql-go/graphql"

var searchBugOrFishResObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "SearchBugOrFishResult",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					// if entry exists
					return entry.Name, nil
				}
				return nil, nil
			},
		},
		"SellPrice": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.SellPrice, nil
				}
				return nil, nil
			},
		},
		"Location": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Location, nil
				}
				return nil, nil
			},
		},
		"Time": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Time, nil
				}
				return nil, nil
			},
		},
		"Type": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Type, nil
				}
				return nil, nil
			},
		},
	},
})
