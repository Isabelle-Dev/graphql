package searchbugandfish

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var listingsObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListingsObject",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if name, ok := p.Source.(string); ok {
					return name, nil
				}
				return nil, fmt.Errorf("error parsing in listings()")
			},
		},
	},
})
