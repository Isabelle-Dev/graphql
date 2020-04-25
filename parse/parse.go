package parse

import (
	"regexp"
	"strings"
)

// QueryParse returns a map of query parsed parameters
//
// E.g. name:"Bob"
//
// Will return a map where the key is 'name'
// and the value is "Bob"
func QueryParse(query string) map[string]string {
	indexes := extractQuoteIndexes(query)
	if len(indexes)%2 != 0 {
		return nil
	}
	fields := extractFields(indexes, query)
	q := regexp.MustCompile("\\s+").Split(query, -1)
	p := extractNames(q)
	if len(p) != (len(indexes) / 2) {
		// there must be matching open and end quotes for every name
		return nil
	}
	options := make(map[string]string, 0)
	for i := 0; i < len(fields); i++ {
		options[p[i]] = fields[i]
	}
	return options
}

// extractFields returns the field inside the query
func extractFields(indexes []int, query string) []string {
	var ret []string
	for i := 0; i < len(indexes); i += 2 {
		ret = append(ret, query[indexes[i]+1:indexes[i+1]])
	}
	return ret
}

// extractName returns a slice of strings that should be names to query options
func extractName(str []string) []string {
	var ret []string
	for _, i := range str {
		index := strings.Index(i, ":")
		ret = append(ret, i[:index])
	}
	return ret
}

// extractNamePotentials returns all strings which the name needs to be extracted from
func extractNames(str []string) []string {
	var ret []string
	for _, s := range str {
		if strings.Index(s, ":") == -1 {
			continue
		}
		ret = append(ret, s)
	}
	return extractName(ret)
}

// extractQuoteIndexes returns an int slice of where all quote indexes are in the string
func extractQuoteIndexes(query string) []int {
	i := strings.Index(query, "'")
	if i == -1 {
		return nil
	}
	var indexes []int
	var prevIndex int
	indexes = append(indexes, i)
	prevIndex = i
	i++
	for i < len(query)-1 {
		i = strings.Index(query[i:], "'")
		if i == -1 {
			return indexes
		}
		i += prevIndex + 1
		indexes = append(indexes, i)
		prevIndex = i
		i++
	}
	return indexes
}
