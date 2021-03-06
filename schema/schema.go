package schema

import (
	"fmt"

	"github.com/Isabelle-Dev/graphql/schema/art"
	"github.com/Isabelle-Dev/graphql/schema/bugs"
	"github.com/Isabelle-Dev/graphql/schema/clothes"
	"github.com/Isabelle-Dev/graphql/schema/construction"
	"github.com/Isabelle-Dev/graphql/schema/fencing"
	"github.com/Isabelle-Dev/graphql/schema/fish"
	"github.com/Isabelle-Dev/graphql/schema/floor"
	"github.com/Isabelle-Dev/graphql/schema/fossil"
	"github.com/Isabelle-Dev/graphql/schema/item"
	"github.com/Isabelle-Dev/graphql/schema/music"
	"github.com/Isabelle-Dev/graphql/schema/nookmiles"
	"github.com/Isabelle-Dev/graphql/schema/other"
	"github.com/Isabelle-Dev/graphql/schema/photos"
	"github.com/Isabelle-Dev/graphql/schema/poster"
	"github.com/Isabelle-Dev/graphql/schema/reaction"
	rug "github.com/Isabelle-Dev/graphql/schema/rugs"
	"github.com/Isabelle-Dev/graphql/schema/tool"
	"github.com/Isabelle-Dev/graphql/schema/umbrella"
	"github.com/Isabelle-Dev/graphql/schema/villager"
	"github.com/Isabelle-Dev/graphql/schema/wallpaper"
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
				Description: "Rug-related queries",
				Type:        rug.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"photos": &graphql.Field{
				Description: "Photo-related queries",
				Type:        photos.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"villagers": &graphql.Field{
				Description: "Villager-related queries",
				Type:        villager.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"art": &graphql.Field{
				Description: "Art-related queries",
				Type:        art.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"poster": &graphql.Field{
				Description: "Poster-related queries",
				Type:        poster.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"music": &graphql.Field{
				Description: "Music-related queries",
				Type:        music.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"tool": &graphql.Field{
				Description: "Tool-related queries",
				Type:        tool.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"construction": &graphql.Field{
				Description: "Construction-related queries",
				Type:        construction.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"nookmiles": &graphql.Field{
				Description: "NookMile-related queries",
				Type:        nookmiles.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"fencing": &graphql.Field{
				Description: "Fencing-related queries",
				Type:        fencing.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"umbrella": &graphql.Field{
				Description: "Umbrella-related queries",
				Type:        umbrella.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"other": &graphql.Field{
				Description: "Other-related queries",
				Type:        other.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"bugs": &graphql.Field{
				Description: "Bug-related queries",
				Type:        bugs.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"fishes": &graphql.Field{
				Description: "Fish-related queries",
				Type:        fish.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"fossil": &graphql.Field{
				Description: "Fossil-related queries",
				Type:        fossil.RootObject(db),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"reaction": &graphql.Field{
				Description: "Reaction-related queries",
				Type:        reaction.RootObject(db),
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

// provides the dialect type for gorm configuration
func (pc PostgresConfig) dialect() string {
	return "postgres"
}

// provides the connection info for gorm configuration
func (pc PostgresConfig) connectionInfo() string {
	if pc.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s "+
			"sslmode=disable", pc.Host, pc.Port, pc.User, pc.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", pc.Host, pc.Port, pc.User, pc.Password, pc.Name)
}
