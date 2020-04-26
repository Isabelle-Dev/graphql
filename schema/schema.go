package schema

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/schema/searchclothes"
	"github.com/Isabelle-Dev/isabelle-graphql/schema/searchitem"
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
			// 	Type:        searchbugandfish.RootObject(db),
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return map[string]interface{}{}, nil
			// 	},
			// },
			"item": &graphql.Field{
				Description: "Item-related queries",
				Type:        searchitem.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"clothes": &graphql.Field{
				Description: "Clothes-related queries",
				Type:        searchclothes.RootObject(db),
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
