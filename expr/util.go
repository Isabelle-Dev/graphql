package expr

import (
	"strconv"
	"strings"
)

// utility func which places parenthesis around given string
func paren(str string) string {
	return "(" + str + ")"
}

// utility func to add the query type's name to query string
func addName(str string, info Env) string {
	switch info["name"] {
	case "concept":
		return "(hhaconcept1 = " + str + " OR hhaconcept2 = " + str + ")"
	case "color":
		return "(color1 = " + str + " OR color2 = " + str + ")"
	case "tag":
		return info["name"] + " LIKE '%" + strings.Trim(str, "'") + "%'"
	case "name":
		if strings.ToLower(info["glob"]) == "t" {
			return info["name"] + " LIKE '%" + strings.Trim(str, "'") + "%'"
		}
		return info["name"] + " = " + str
	case "theme":
		return "labelthemes LIKE '%" + strings.Trim(str, "'") + "%'"
	case "style":
		return "(style1 = " + str + " OR style2 = " + str + ")"
	default:
		return info["name"] + " = " + str
	}
}

// utility func to add quotes around string literals
func addQuote(str string) string {
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.TrimSpace(str)
	if _, err := strconv.Atoi(str); err == nil {
		// we don't add quotes around numbers
		return str
	}
	return "'" + str + "'"
}
