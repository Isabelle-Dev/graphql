package rug

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// Converts a slice of RugEntry into a slice of *Rug
func toRugSlice(rugs []newhorizons.RugEntry) []*newhorizons.Rug {
	var ret []*newhorizons.Rug
	for _, i := range rugs {
		color, concept := extractCC(i)
		ret = append(ret, i.ToGraphQL(color, concept))
	}
	return ret
}

// Extract and combine categories (color and concepts)
func extractCC(entry newhorizons.RugEntry) ([]string, []string) {
	var color []string
	var concept []string
	color = append(color, entry.Color1)
	if !common.Exists(entry.Color2, color) {
		color = append(color, entry.Color2)
	}
	concept = append(concept, entry.HHAConcept1)
	if !common.Exists(entry.HHAConcept2, concept) && entry.HHAConcept2 != "None" {
		concept = append(concept, entry.HHAConcept2)
	}
	return color, concept
}
