package search

import "github.com/graphql-go/graphql"

var searchBugOrFishResObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "SearchBugOrFishResult",
	Fields: graphql.Fields{
		"Num": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Num, nil
				}
				return nil, nil
			},
		},
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Image, nil
				}
				return nil, nil
			},
		},
		"House": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.House, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Sell, nil
				}
				return nil, nil
			},
		},
		"Where": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Where, nil
				}
				return nil, nil
			},
		},
		"Weather": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Weather, nil
				}
				return nil, nil
			},
		},
		"Rarity": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Rarity, nil
				}
				return nil, nil
			},
		},
		"StartTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.StartTime, nil
				}
				return nil, nil
			},
		},
		"EndTime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.EndTime, nil
				}
				return nil, nil
			},
		},
		"Jan": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Jan, nil
				}
				return nil, nil
			},
		},
		"Feb": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Feb, nil
				}
				return nil, nil
			},
		},
		"Mar": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Mar, nil
				}
				return nil, nil
			},
		},
		"Apr": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Apr, nil
				}
				return nil, nil
			},
		},
		"May": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.May, nil
				}
				return nil, nil
			},
		},
		"Jun": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Jun, nil
				}
				return nil, nil
			},
		},
		"Jul": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Jul, nil
				}
				return nil, nil
			},
		},
		"Aug": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Aug, nil
				}
				return nil, nil
			},
		},
		"Sep": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Sep, nil
				}
				return nil, nil
			},
		},
		"Oct": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Oct, nil
				}
				return nil, nil
			},
		},
		"Nov": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Nov, nil
				}
				return nil, nil
			},
		},
		"Dec": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Dec, nil
				}
				return nil, nil
			},
		},
		"Color1": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Color1, nil
				}
				return nil, nil
			},
		},
		"Color2": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.Color2, nil
				}
				return nil, nil
			},
		},
		"ItemFilename": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.ItemFilename, nil
				}
				return nil, nil
			},
		},
		"InternalID": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entry, ok := p.Source.(*Entry); ok {
					return entry.InternalID, nil
				}
				return nil, nil
			},
		},
	},
})
