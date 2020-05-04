package fish

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// Convert a slice of FishEntry into a slice of *Fish
func toFishSlice(f []newhorizons.FishEntry) []*newhorizons.Fish {
	var ret []*newhorizons.Fish
	for _, entry := range f {
		c := extractC(entry)
		ret = append(ret, entry.ToGraphQL(c))
	}
	return ret
}

// Extract the color from entry as long as it's not a duplicate
//
// Color1 is always guaranteed a color for fishes
func extractC(f newhorizons.FishEntry) []string {
	var ret []string
	ret = append(ret, f.Color1)
	if !common.Exists(f.Color2, ret) {
		ret = append(ret, f.Color2)
	}
	return ret
}
