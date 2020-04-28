package tool

import (
	"strings"

	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func toToolSlice(t []newhorizons.ToolEntry, db *gorm.DB) []*newhorizons.Tool {
	dupe := make(map[string]bool, 0)
	var ret []*newhorizons.Tool
	for _, i := range t {
		if _, ok := dupe[i.Name]; ok {
			continue
		}
		n := findByName(i.Name, "tool", db)
		dupe[i.Name] = true
		ret = append(ret, i.ToGraphQL(n.Variant, n.Source))
	}
	return ret
}

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
