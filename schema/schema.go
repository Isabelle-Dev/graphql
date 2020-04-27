package schema

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/schema/clothes"
	"github.com/Isabelle-Dev/isabelle-graphql/schema/floor"
	"github.com/Isabelle-Dev/isabelle-graphql/schema/item"
	"github.com/Isabelle-Dev/isabelle-graphql/schema/photos"
	rug "github.com/Isabelle-Dev/isabelle-graphql/schema/rugs"
	"github.com/Isabelle-Dev/isabelle-graphql/schema/wallpaper"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	// need this import for postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Root of schema
var Root graphql.Schema

func init() {
	// set postgres config
	cfg, er := loadConfig()
	if er != nil {
		panic(er)
	}
	// initialize database connection
	db, e := gorm.Open(cfg.Database.dialect(), cfg.Database.connectionInfo())
	if e != nil {
		panic(e)
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Root query",
		Fields: graphql.Fields{
			// "bug_and_fish": &graphql.Field{
			// 	Description: "Bug and fish search-related queries",
			// 	Type:        bugandfish.RootObject(db),
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return map[string]interface{}{}, nil
			// 	},
			// },
			"item": &graphql.Field{
				Description: "Item-related queries",
				Type:        item.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"clothes": &graphql.Field{
				Description: "Clothes-related queries",
				Type:        clothes.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"wallpaper": &graphql.Field{
				Description: "Wallpaper-related queries",
				Type:        wallpaper.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"floor": &graphql.Field{
				Description: "Floor-related queries",
				Type:        floor.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"rugs": &graphql.Field{
				Description: "Floor-related queries",
				Type:        rug.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"photos": &graphql.Field{
				Description: "Floor-related queries",
				Type:        photos.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
		},
	})

	var err error
	Root, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		panic(err)
	}
}

func (pc PostgresConfig) dialect() string {
	return "postgres"
}

func (pc PostgresConfig) connectionInfo() string {
	if pc.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s "+
			"sslmode=disable", pc.Host, pc.Port, pc.User, pc.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", pc.Host, pc.Port, pc.User, pc.Password, pc.Name)
}
