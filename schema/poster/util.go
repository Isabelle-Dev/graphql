package poster

import (
	"github.com/Isabelle-Dev/isabelle-graphql/common"
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
)

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
