package music

import (
	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// Converts a slice of MusicEntry into a slice of *Music
func toMusicSlice(m []newhorizons.MusicEntry) []*newhorizons.Music {
	var ret []*newhorizons.Music
	for _, i := range m {
		c := extractC(i)
		ret = append(ret, i.ToGraphQL(c))
	}
	return ret
}

// Extract and combine unique color info from a MusicEntry
func extractC(m newhorizons.MusicEntry) []string {
	var ret []string
	ret = append(ret, m.Color1)
	if !common.Exists(m.Color2, ret) && m.Color2 != "None" {
		ret = append(ret, m.Color2)
	}
	return ret
}
