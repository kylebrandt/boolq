package parse

import (
	"testing"
)

type parseTest struct {
	name   string
	input  string
	ok     bool
	result string // what the user would see in an error message.
}

const (
	noError  = true
	hasError = false
)

var parseTests = []parseTest{
	{"single ask", "ask:something", noError, ""},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		q, err := Parse(test.input)
		if err != nil {
			t.Error(err)
		}
		_ = q
	}
}
