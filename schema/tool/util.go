package tool

import (
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
		ret = append(ret, i.ToGraphQL(n.Variant))
	}
	return ret
}

func extractV(t []newhorizons.ToolEntry) []newhorizons.ToolVariant {
	var v []newhorizons.ToolVariant
	for _, entry := range t {
		v = append(v, newhorizons.ToolVariant{
			Image:     entry.Image,
			Variation: entry.Variation,
		})
	}
	return v
}
