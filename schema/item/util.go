package item

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// helper func which extracts all variant, pattern, images, and HHA concept tags
func extractVPIH(item []newhorizons.ItemEntry) ([]newhorizons.Variant, []string) {
	var variants []newhorizons.Variant
	var hha []string
	for _, entry := range item {
		// Add variant details - remove duplicates
		var color []string
		if !exists(entry.Color1, color) {
			color = append(color, entry.Color1)
		}
		if !exists(entry.Color2, color) {
			color = append(color, entry.Color2)
		}
		variants = append(variants, newhorizons.Variant{
			ImageURL: entry.Image,
			Pattern:  entry.Pattern,
			Colors:   color,
		})
	}
	// extract HHA concepts
	if !exists(item[0].HHAConcept1, hha) && item[0].HHAConcept1 != "None" {
		hha = append(hha, item[0].HHAConcept1)
	}
	if !exists(item[0].HHAConcept2, hha) && item[0].HHAConcept2 != "None" {
		hha = append(hha, item[0].HHAConcept2)
	}
	return variants, hha
}

// utility func to check existence
func exists(match string, container []string) bool {
	for _, i := range container {
		if match == i {
			return true
		}
	}
	return false
}

// utility func to turn all ItemEntry entries into Item
func toItemSlice(items []newhorizons.ItemEntry, db *gorm.DB) []*newhorizons.Item {
	var names []string
	var ret []*newhorizons.Item
	for _, i := range items {
		if exists(i.Name, names) {
			continue
		}
		n := findByName(i.Name, "item", db)
		names = append(names, i.Name)
		ret = append(ret, i.ToGraphQL(n.Variants, n.HHAConcepts))
	}
	return ret
}
