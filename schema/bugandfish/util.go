package bugandfish

import (
	"reflect"

	"github.com/Isabelle-Dev/graphql/common"
)

// helper func which extracts month data from an entry
func extractMonths(entry interface{}) []string {
	months := []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}
	var ret []string
	v := reflect.ValueOf(entry)
	vtype := reflect.TypeOf(entry)
	for i := 0; i < v.NumField(); i++ {
		if common.Exists(vtype.Field(i).Name, months) {
			// is a month type
			if v.Field(i).String() == "Yes" {
				ret = append(ret, vtype.Field(i).Name)
			}
		}
	}
	return ret
}
