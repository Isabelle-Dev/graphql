package bugs

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// convert a slice of BugEntry into a slice of *Bug
func toBugSlice(b []newhorizons.BugEntry) []*newhorizons.Bug {
	var ret []*newhorizons.Bug
	for _, entry := range b {
		c := extractC(entry)
		ret = append(ret, entry.ToGraphQL(c))
	}
	return ret
}

// extract the color from entry as long as it's not a duplicate
//
// color1 is always guaranteed a color for bugs
func extractC(entry newhorizons.BugEntry) []string {
	var ret []string
	ret = append(ret, entry.Color1)
	if !common.Exists(entry.Color2, ret) {
		ret = append(ret, entry.Color2)
	}
	return ret
}
