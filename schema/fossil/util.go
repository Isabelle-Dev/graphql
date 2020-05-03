package fossil

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// toFossilSlice returns a slice of Fossil objects converted from FossilEntry
func toFossilSlice(f []newhorizons.FossilEntry) []*newhorizons.Fossil {
	var ret []*newhorizons.Fossil
	for _, entry := range f {
		c := extractC(entry)
		ret = append(ret, entry.ToGraphQL(c))
	}
	return ret
}

// extractC returns a slice of colors that an entry has without duplicates
func extractC(entry newhorizons.FossilEntry) []string {
	var ret []string
	ret = append(ret, entry.Color1)
	if !common.Exists(entry.Color2, ret) {
		ret = append(ret, entry.Color2)
	}
	return ret
}
