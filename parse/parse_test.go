package parse

import (
	"testing"
)

func TestParseQuery(t *testing.T) {
	tests := map[string]struct {
		input     string
		wantType  string
		wantField string
	}{
		"simple case": {
			input:     "name:\"Bob\"",
			wantType:  "name",
			wantField: "Bob",
		},
		"multiple": {
			input:     "one:\"here\" another:\"lol\" thisshouldn'tshow",
			wantType:  "meh",
			wantField: "meh",
		},
		"check three": {
			input:     "one:\"lol\" two:\"haha\" three:\"Ciao ha\"",
			wantType:  "eh",
			wantField: "eh",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := QueryParse(tc.input)
			if _, ok := got[tc.wantType]; !ok {
				t.Errorf("parseQuery(): error in wantType; want %s", tc.wantType)
			}
			if got[tc.wantType] != tc.wantField {
				t.Errorf("parseQuery(): error; got %s; want %s", got[tc.wantType], tc.wantField)
			}
		})
	}
}

func TestExtractQuoteIndexes(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []int
	}{
		"simple case": {
			input: "hello:\"right\"",
			want:  []int{6, 12},
		},
		"multiple": {
			input: "one:\"here\" another:\"lol\"",
			want:  []int{4, 9, 19, 23},
		},
		"check three": {
			input: "one:\"lol\" two:\"haha\" three:\"Ciao ha\"",
			want:  []int{4, 8, 14, 19, 27, 35},
		},
		"error nil": {
			input: "nothing",
			want:  nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := extractQuoteIndexes(tc.input)
			if got == nil {
				t.Errorf("something went wrong")
			}
			for _, num := range got {
				if !contains(num, tc.want) {
					t.Errorf("extractQuoteIndexes(): got %d; want none", num)
				}
			}
		})
	}
}

func contains(num int, nums []int) bool {
	for _, i := range nums {
		if i == num {
			return true
		}
	}
	return false
}
