package common

import "github.com/Isabelle-Dev/graphql/newhorizons"

// Exists is a utility func which returns true or false if a string slice contains a
// specified string element
func Exists(element string, container []string) bool {
	for _, i := range container {
		if i == element {
			return true
		}
	}
	return false
}

// ToMonthSlice is a utility func which compiles all new horizons month types into a single slice
func ToMonthSlice(m ...newhorizons.Month) []newhorizons.Month { return m }
