package poster

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// Converts a slice of PosterEntry into a slice of *Poster
func toPosterSlice(p []newhorizons.PosterEntry) []*newhorizons.Poster {
	var ret []*newhorizons.Poster
	for _, i := range p {
		var color []string
		color = append(color, i.Color1)
		if !common.Exists(i.Color2, color) {
			color = append(color, i.Color2)
		}
		ret = append(ret, i.ToGraphQL(color))
	}
	return ret
}
