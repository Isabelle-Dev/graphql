package searchclothes

import (
	"strings"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func toClothesSlice(clothes []newhorizons.ClothesEntry, db *gorm.DB) []*newhorizons.Clothes {
	var names []string
	var ret []*newhorizons.Clothes
	for _, c := range clothes {
		if exists(c.Name, names) {
			continue
		}
		n := findByName(c.Name, "clothes", db)
		names = append(names, c.Name)
		ret = append(ret, c.ToGraphQL(n.Variants, n.Themes))
	}
	return ret
}

func extractVT(c []newhorizons.ClothesEntry) ([]newhorizons.ClothesVariant, []string) {
	var v []newhorizons.ClothesVariant
	for _, entry := range c {
		// extract unique variant data
		var color []string
		if !exists(entry.Color1, color) {
			color = append(color, entry.Color1)
		}
		if !exists(entry.Color2, color) {
			color = append(color, entry.Color2)
		}
		v = append(v, newhorizons.ClothesVariant{
			Image:     entry.Image,
			Variation: entry.Variation,
			Color:     color,
		})
	}
	// extract themes
	t := c[0].Themes
	var themes []string
	ts := strings.Split(t, ";")
	for _, w := range ts {
		word := strings.TrimSpace(w)
		if exists(word, themes) {
			continue
		}
		themes = append(themes, word)
	}
	return v, themes
}

func exists(c string, i []string) bool {
	for _, k := range i {
		if k == c {
			return true
		}
	}
	return false
}
