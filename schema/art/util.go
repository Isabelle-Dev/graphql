package art

import (
	"github.com/Isabelle-Dev/isabelle-graphql/common"
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
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
		ret = append(ret, i.ToGraphQL(n.Type, n.Concept))
	}
	return ret
}

func extractType(a []newhorizons.ArtEntry) ([]newhorizons.ArtType, []string) {
	var at []newhorizons.ArtType
	for _, entry := range a {
		at = append(at, newhorizons.ArtType{
			Image:   entry.Image,
			Sell:    entry.Sell,
			Genuine: entry.Genuine,
		})
	}
	var concept []string
	concept = append(concept, a[0].HHAConcept1)
	if !common.Exists(a[0].HHAConcept2, concept) {
		concept = append(concept, a[0].HHAConcept2)
	}
	return at, concept
}
