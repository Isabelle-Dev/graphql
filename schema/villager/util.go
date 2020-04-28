package villager

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

func toVillagerSlice(v []newhorizons.VillagerEntry) []*newhorizons.Villager {
	var ret []*newhorizons.Villager
	for _, i := range v {
		s, c := extractSC(i)
		ret = append(ret, i.ToGraphQL(s, c))
	}
	return ret
}

func extractSC(entry newhorizons.VillagerEntry) ([]string, []string) {
	var style []string
	var color []string
	style = append(style, entry.Style1)
	if !common.Exists(entry.Style2, style) {
		style = append(style, entry.Style2)
	}
	color = append(color, entry.Color1)
	if !common.Exists(entry.Color2, color) {
		color = append(color, entry.Color2)
	}
	return style, color
}
