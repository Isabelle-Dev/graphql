package clothes

import (
	"strings"

	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// Convert a slice of ClothesEntry into a slice of *Clothes
func toClothesSlice(clothes []newhorizons.ClothesEntry, db *gorm.DB) []*newhorizons.Clothes {
	dupe := make(map[string]bool, 0)
	var ret []*newhorizons.Clothes
	for _, c := range clothes {
		if _, ok := dupe[c.Name]; ok {
			continue
		}
		n := findByName(c.Name, "clothes", db)
		dupe[c.Name] = true
		ret = append(ret, c.ToGraphQL(n.Variants, n.Themes))
	}
	return ret
}

// Extract color, ClothesVariant, and theme data from a slice of ClothesEntry.
//
// This is mainly used to combine variant data into a singular object
func extractVT(c []newhorizons.ClothesEntry) ([]newhorizons.ClothesVariant, []string) {
	var v []newhorizons.ClothesVariant
	for _, entry := range c {
		// extract unique variant data
		var color []string
		color = append(color, entry.Color1)
		if !common.Exists(entry.Color2, color) {
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
		if common.Exists(word, themes) {
			continue
		}
		themes = append(themes, word)
	}
	return v, themes
}
