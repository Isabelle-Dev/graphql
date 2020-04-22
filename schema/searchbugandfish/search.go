package searchbugandfish

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var bugAndFishSearchObject *graphql.Object

// RootObject that contains search related queries
func RootObject(db *gorm.DB) *graphql.Object {
	if bugAndFishSearchObject != nil {
		return bugAndFishSearchObject
	}

	// Search
	bugAndFishSearchObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "bug_and_fish",
		Description: "Bug and Fish-related Sources",
		Fields: graphql.Fields{
			"search_bug": &graphql.Field{
				Name: "SearchABugFromNorthernHemisphere",
				Type: graphql.NewNonNull(searchBugResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)
					entry := findBugByName(item, "north_bug", "south_bug", db)
					if entry == nil {
						return nil, fmt.Errorf("search(): %v not found", item)
					}
					return entry, nil
				},
			},
			"search_fish": &graphql.Field{
				Name: "SearchAFishFromNorthernHemisphere",
				Type: graphql.NewNonNull(searchFishResObj),
				Args: graphql.FieldConfigArgument{
					"item": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					item := p.Args["item"].(string)

					entry := findFishByName(item, "north_fish", "south_fish", db)
					if entry == nil {
						return nil, fmt.Errorf("search(): %v not found", item)
					}
					return entry, nil
				},
			},

			"search_all_by_hemisphere": &graphql.Field{
				Name: "SearchAllBugAndFishInTheNorth",
				Type: graphql.NewNonNull(searchCombinedObj),
				Args: graphql.FieldConfigArgument{
					"hemi": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					hemi := p.Args["hemi"].(string)
					if hemi == "north" {
						entries := findAllByHemisphere("north_bug", "north_fish", db)
						if entries == nil {
							return nil, fmt.Errorf("search() by hemi: Error in database connection")
						}
						return entries, nil
					} else if hemi == "south" {
						entries := findAllByHemisphere("south_bug", "south_fish", db)
						if entries == nil {
							return nil, fmt.Errorf("search() by hemi: Error in database connection")
						}
						return entries, nil
					} else {
						return nil, fmt.Errorf("search(%v): invalid hemisphere search", hemi)
					}
				},
			},

			"search_by_price": &graphql.Field{
				Name: "SearchBugOrFishNamesBySellPrice",
				Type: graphql.NewNonNull(searchAgnosticObj),
				Args: graphql.FieldConfigArgument{
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					price := p.Args["price"].(int)
					entries := findByPrice(price, "north_bug", "north_fish", db)
					if entries == nil {
						return nil, fmt.Errorf("search() by price: Error in database connection")
					}
					return entries, nil
				},
			},

			"search_by_month": &graphql.Field{
				Name: "SearchBugAndFishesByMonth",
				Type: graphql.NewNonNull(searchCombinedObj),
				Args: graphql.FieldConfigArgument{
					"month": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"hemi": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					month := p.Args["month"].(string)
					hemi := p.Args["hemi"].(string)
					if hemi == "north" {
						entries := findByMonth(month, "north_bug", "north_fish", db)
						if entries == nil {
							return nil, fmt.Errorf("search(): couldn't create month entries")
						}
						return entries, nil
					} else if hemi == "south" {
						entries := findByMonth(month, "south_bug", "south_fish", db)
						if entries == nil {
							return nil, fmt.Errorf("search(): couldn't create month entries")
						}
						return entries, nil
					} else {
						return nil, fmt.Errorf("search(): invalid hemi in search by month")
					}
				},
			},

			"search_by_rarity": &graphql.Field{
				Name: "SearchBugsAndFishesByRarity",
				Type: graphql.NewNonNull(searchAgnosticObj),
				Args: graphql.FieldConfigArgument{
					"rarity": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					rarity := p.Args["rarity"].(string)
					entries := findByRarity(rarity, "north_bug", "north_fish", db)
					if entries == nil {
						return nil, fmt.Errorf("search(): rarity search failed")
					}
					return entries, nil
				},
			},

			"search_by_shadow": &graphql.Field{
				Name: "SearchAFishByShadowSize",
				Type: graphql.NewNonNull(graphql.NewList(agCombinedFishObj)),
				Args: graphql.FieldConfigArgument{
					"shadow": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					shadow := p.Args["shadow"].(string)
					entries := findByShadow(shadow, "north_fish", db)
					if entries == nil {
						return nil, fmt.Errorf("search(): shadow search failed")
					}
					return entries, nil
				},
			},
		},
	})
	return bugAndFishSearchObject
}
