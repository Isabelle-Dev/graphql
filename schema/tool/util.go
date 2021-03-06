package tool

import (
	"strings"

	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// Convert a slice of ToolEntry into a slice of *Tool
func toToolSlice(t []newhorizons.ToolEntry, db *gorm.DB) []*newhorizons.Tool {
	dupe := make(map[string]bool, 0)
	var ret []*newhorizons.Tool
	for _, i := range t {
		if _, ok := dupe[i.Name]; ok {
			continue
		}
		n := findByName(i.Name, "tools", db)
		dupe[i.Name] = true
		ret = append(ret, i.ToGraphQL(n.Variant, n.Source))
	}
	return ret
}

// Extract and combine variant data into a singular category
func extractVS(t []newhorizons.ToolEntry) ([]newhorizons.ToolVariant, []string) {
	var v []newhorizons.ToolVariant
	for _, entry := range t {
		v = append(v, newhorizons.ToolVariant{
			Image:     entry.Image,
			Variation: entry.Variation,
		})
	}
	var source []string
	parts := strings.Split(t[0].Source, ";")
	for _, i := range parts {
		word := strings.TrimSpace(i)
		source = append(source, word)
	}
	return v, source
}
