package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// helper func which extracts all variant, pattern, images, and HHA concept tags
func extractVPIH(item []newhorizons.ItemEntry) ([]string, []string, []string, []string) {
	var variants []string
	var patterns []string
	var images []string
	var hha []string
	for _, entry := range item {
		if !exists(entry.HHAConcept1, hha) && entry.HHAConcept1 != "None" {
			hha = append(hha, entry.HHAConcept1)
		}
		if !exists(entry.HHAConcept2, hha) && entry.HHAConcept2 != "None" {
			hha = append(hha, entry.HHAConcept2)
		}
		if !exists(entry.PatternTitle, patterns) && entry.PatternTitle != "NA" {
			patterns = append(patterns, entry.Pattern)
		}
		if !exists(entry.Variation, variants) && entry.Variation != "NA" {
			variants = append(variants, entry.Variation)
		}
		images = append(images, entry.Image)
	}
	return variants, patterns, images, hha
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
		ret = append(ret, i.ToGraphQL(n.Variation, n.Pattern, n.Image, n.HHAConcepts))
	}
	return ret
}
