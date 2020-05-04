package parse

import (
	"fmt"
	"strconv"

	"github.com/Isabelle-Dev/graphql/expr"
)

// BuildQuery builds a SQL query string from separated user query returned in
// QueryParse
func BuildQuery(vals map[string]string, tablename string, limit int, globSearch string) string {
	lim := strconv.Itoa(limit)
	if len(vals) == 0 {
		return "SELECT * FROM " + tablename + " LIMIT " + lim
	}
	res := "SELECT * FROM " + tablename + " WHERE"
	first := true
	env := make(map[string]string, 0)
	env["glob"] = globSearch
	for k, v := range vals {
		env["name"] = k
		tmp, err := expr.ParseAndEval(v, expr.Env(env))
		if err != nil {
			// debugging purposes
			fmt.Println(err)
		}
		if first {
			res += " " + tmp
			first = false
			continue
		}
		res += " AND " + tmp
	}
	res += " LIMIT " + lim
	return res
}
