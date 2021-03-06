package umbrella

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// Convert a slice of UmbrellaEntry into a slice of *Umbrella
func toUmbrellaSlice(u []newhorizons.UmbrellaEntry) []*newhorizons.Umbrella {
	var ret []*newhorizons.Umbrella
	for _, i := range u {
		color := extractC(i)
		ret = append(ret, i.ToGraphQL(color))
	}
	return ret
}

// Extract unique color data into a single category
func extractC(u newhorizons.UmbrellaEntry) []string {
	var ret []string
	ret = append(ret, u.Color1)
	if !common.Exists(u.Color2, ret) && u.Color2 != "None" {
		ret = append(ret, u.Color2)
	}
	return ret
}
