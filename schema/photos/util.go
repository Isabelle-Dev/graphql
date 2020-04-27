package photos

import (
	"github.com/Isabelle-Dev/isabelle-graphql/common"
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func toPhotoSlice(p []newhorizons.PhotoEntry, db *gorm.DB) []*newhorizons.Photo {
	dupe := make(map[string]bool)
	var ret []*newhorizons.Photo
	for _, i := range p {
		if _, ok := dupe[i.Name]; ok {
			continue
		}
		n := findByName(i.Name, "photos", db)
		dupe[i.Name] = true
		ret = append(ret, i.ToGraphQL(n.Variants))
	}
	return ret
}

func extractV(p []newhorizons.PhotoEntry) []newhorizons.Variant {
	var v []newhorizons.Variant
	for _, entry := range p {
		var color []string
		color = append(color, entry.Color1)
		if !common.Exists(entry.Color2, color) {
			color = append(color, entry.Color2)
		}
		v = append(v, newhorizons.Variant{
			ImageURL: entry.Image,
			Pattern:  entry.Variation,
			Colors:   color,
		})
	}
	return v
}
