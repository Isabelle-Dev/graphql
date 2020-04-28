package floor

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

func toFloorSlice(floors []newhorizons.FloorEntry) []*newhorizons.Floor {
	var ret []*newhorizons.Floor
	for _, i := range floors {
		color, concept := buildSchema(i)
		ret = append(ret, i.ToGraphQL(color, concept))
	}
	return ret
}

func buildSchema(entry newhorizons.FloorEntry) ([]string, []string) {
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
