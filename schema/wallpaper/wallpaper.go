package wallpaper

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/graphql-go/graphql"
)

var wallpaper = graphql.NewObject(graphql.ObjectConfig{
	Name: "WallpaperResult",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Name, nil
				}
				return nil, nil
			},
		},
		"Image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Image, nil
				}
				return nil, nil
			},
		},
		"VFXInfo": &graphql.Field{
			Type: vfxObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.VFXInfo, nil
				}
				return nil, nil
			},
		},
		"DIY": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.DIY, nil
				}
				return nil, nil
			},
		},
		"Buy": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Buy, nil
				}
				return nil, nil
			},
		},
		"Sell": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Sell, nil
				}
				return nil, nil
			},
		},
		"Source": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Source, nil
				}
				return nil, nil
			},
		},
		"SourceNotes": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.SourceNotes, nil
				}
				return nil, nil
			},
		},
		"WindowInfo": &graphql.Field{
			Type: windowObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.WindowInfo, nil
				}
				return nil, nil
			},
		},
		"CurtainInfo": &graphql.Field{
			Type: curtainObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.CurtainInfo, nil
				}
				return nil, nil
			},
		},
		"CeilingType": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.CeilingType, nil
				}
				return nil, nil
			},
		},
		"Colors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Color, nil
				}
				return nil, nil
			},
		},
		"Concepts": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Concepts, nil
				}
				return nil, nil
			},
		},
		"HHASeries": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.HHASeries, nil
				}
				return nil, nil
			},
		},
		"Tag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*newhorizons.Wallpaper); ok {
					return val.Tag, nil
				}
				return nil, nil
			},
		},
	},
})

var vfxObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "VFX",
	Fields: graphql.Fields{
		"VFX": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.VFXT); ok {
					return val.VFX, nil
				}
				return nil, nil
			},
		},
		"VFXType": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.VFXT); ok {
					return val.VFXType, nil
				}
				return nil, nil
			},
		},
	},
})

var windowObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Window",
	Fields: graphql.Fields{
		"Type": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Windows); ok {
					return val.WindowType, nil
				}
				return nil, nil
			},
		},
		"Color": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Windows); ok {
					return val.WindowColor, nil
				}
				return nil, nil
			},
		},
	},
})

var curtainObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Curtain",
	Fields: graphql.Fields{
		"CurtainType": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Curtains); ok {
					return val.CurtainType, nil
				}
				return nil, nil
			},
		},
		"CurtainColor": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(newhorizons.Curtains); ok {
					return val.CurtainColor, nil
				}
				return nil, nil
			},
		},
	},
})
