package fish

import (
	"strings"

	"github.com/Isabelle-Dev/graphql/common"
	"github.com/Isabelle-Dev/graphql/newhorizons"
)

// Convert a slice of FishEntry into a slice of *Fish
func toFishSlice(f []newhorizons.FishEntry, month string) []*newhorizons.Fish {
	var ret []*newhorizons.Fish
	for _, entry := range f {
		c := extractC(entry)
		m := extractMonths(entry, month)
		ret = append(ret, entry.ToGraphQL(c, m))
	}
	return ret
}

func extractMonths(entry newhorizons.FishEntry, m string) []newhorizons.Month {
	if m == "all" {
		// return all month times
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Jan",
				Time:  entry.NHJan,
			},
			newhorizons.Month{
				Month: "NH Feb",
				Time:  entry.NHFeb,
			},
			newhorizons.Month{
				Month: "NH Mar",
				Time:  entry.NHMar,
			},
			newhorizons.Month{
				Month: "NH Apr",
				Time:  entry.NHApr,
			},
			newhorizons.Month{
				Month: "NH May",
				Time:  entry.NHMay,
			},
			newhorizons.Month{
				Month: "NH Jun",
				Time:  entry.NHJun,
			},
			newhorizons.Month{
				Month: "NH Jul",
				Time:  entry.NHJul,
			},
			newhorizons.Month{
				Month: "NH Aug",
				Time:  entry.NHAug,
			},
			newhorizons.Month{
				Month: "NH Sep",
				Time:  entry.NHSep,
			},
			newhorizons.Month{
				Month: "NH Oct",
				Time:  entry.NHOct,
			},
			newhorizons.Month{
				Month: "NH Nov",
				Time:  entry.NHNov,
			},
			newhorizons.Month{
				Month: "NH Dec",
				Time:  entry.NHDec,
			},
			newhorizons.Month{
				Month: "SH Jan",
				Time:  entry.SHJan,
			},
			newhorizons.Month{
				Month: "SH Feb",
				Time:  entry.SHFeb,
			},
			newhorizons.Month{
				Month: "SH Mar",
				Time:  entry.SHMar,
			},
			newhorizons.Month{
				Month: "SH Apr",
				Time:  entry.SHApr,
			},
			newhorizons.Month{
				Month: "SH May",
				Time:  entry.SHMay,
			},
			newhorizons.Month{
				Month: "SH Jun",
				Time:  entry.SHJun,
			},
			newhorizons.Month{
				Month: "SH Jul",
				Time:  entry.SHJul,
			},
			newhorizons.Month{
				Month: "SH Aug",
				Time:  entry.SHAug,
			},
			newhorizons.Month{
				Month: "SH Sep",
				Time:  entry.SHSep,
			},
			newhorizons.Month{
				Month: "SH Oct",
				Time:  entry.SHOct,
			},
			newhorizons.Month{
				Month: "SH Nov",
				Time:  entry.SHNov,
			},
			newhorizons.Month{
				Month: "SH Dec",
				Time:  entry.SHDec,
			},
		)
	}
	return getIndividualMonth(entry, m)
}

func getIndividualMonth(e newhorizons.FishEntry, m string) []newhorizons.Month {
	switch strings.ToLower(m) {
	case "jan":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Jan",
				Time:  e.NHJan,
			},
			newhorizons.Month{
				Month: "SH Jan",
				Time:  e.SHJan,
			},
		)
	case "feb":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Feb",
				Time:  e.NHFeb,
			},
			newhorizons.Month{
				Month: "SH Feb",
				Time:  e.SHFeb,
			},
		)
	case "mar":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Mar",
				Time:  e.NHMar,
			},
			newhorizons.Month{
				Month: "SH Mar",
				Time:  e.SHMar,
			},
		)
	case "apr":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Apr",
				Time:  e.NHApr,
			},
			newhorizons.Month{
				Month: "SH Apr",
				Time:  e.SHApr,
			},
		)
	case "may":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH May",
				Time:  e.NHMay,
			},
			newhorizons.Month{
				Month: "SH May",
				Time:  e.SHMay,
			},
		)
	case "jun":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Jun",
				Time:  e.NHJun,
			},
			newhorizons.Month{
				Month: "SH Jun",
				Time:  e.SHJun,
			},
		)
	case "jul":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Jul",
				Time:  e.NHJul,
			},
			newhorizons.Month{
				Month: "SH Jul",
				Time:  e.SHJul,
			},
		)
	case "aug":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Aug",
				Time:  e.NHAug,
			},
			newhorizons.Month{
				Month: "SH Aug",
				Time:  e.SHAug,
			},
		)
	case "sep":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Sep",
				Time:  e.NHSep,
			},
			newhorizons.Month{
				Month: "SH Sep",
				Time:  e.SHSep,
			},
		)
	case "oct":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Oct",
				Time:  e.NHOct,
			},
			newhorizons.Month{
				Month: "SH Oct",
				Time:  e.SHOct,
			},
		)
	case "nov":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Nov",
				Time:  e.NHNov,
			},
			newhorizons.Month{
				Month: "SH Nov",
				Time:  e.SHNov,
			},
		)
	case "dec":
		return common.ToMonthSlice(
			newhorizons.Month{
				Month: "NH Dec",
				Time:  e.NHDec,
			},
			newhorizons.Month{
				Month: "SH Dec",
				Time:  e.SHDec,
			},
		)
	}
	return nil
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
