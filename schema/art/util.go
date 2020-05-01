package art

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func toArtSlice(a []newhorizons.ArtEntry, db *gorm.DB) []*newhorizons.Art {
	dupe := make(map[string]bool, 0)
	var ret []*newhorizons.Art
	for _, i := range a {
		if _, ok := dupe[i.Name]; ok {
			continue
		}
		n := findByName(i.Name, "art", db)
		dupe[i.Name] = true
		ret = append(ret, i.ToGraphQL(n.Type))
	}
	return ret
}

func extractType(a []newhorizons.ArtEntry) []newhorizons.ArtType {
	var at []newhorizons.ArtType
	for _, entry := range a {
		var concept []string
		concept = append(concept, entry.HHAConcept1)
		if !common.Exists(entry.HHAConcept2, concept) && entry.HHAConcept2 != "None" {
			concept = append(concept, entry.HHAConcept2)
		}
		at = append(at, newhorizons.ArtType{
			Image:   entry.Image,
			Sell:    entry.Sell,
			Genuine: entry.Genuine,
			Concept: concept,
		})
	}
	return at
}
