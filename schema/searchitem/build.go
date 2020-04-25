package searchitem

import (
	"fmt"

	"github.com/Isabelle-Dev/isabelle-graphql/expr"
)

// E.g. name:"Bob"
//
// Will return a map where the key is 'name'
// and the value is "Bob"
func buildQuery(vals map[string]string, tablename string) string {
	if len(vals) == 0 {
		return "SELECT * FROM " + tablename
	}
	res := "SELECT * FROM " + tablename + " WHERE"
	first := true
	for k, v := range vals {
		tmp, err := expr.ParseAndEval(v, expr.Env(k))
		if err != nil {
			fmt.Println(err)
		}
		if first {
			res += " " + tmp
			first = false
			continue
		}
		res += " AND " + tmp
	}
	return res
}
