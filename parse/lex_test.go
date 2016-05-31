package parse

import "testing"

type lexTest struct {
	name  string
	input string
	items []item
}

var (
	tAND = item{itemAnd, 0, "AND"}
	tEOF = item{itemEOF, 0, ""}
)

var lexTests = []lexTest{
	{"empty", "", []item{tEOF}},
	{"and test", "AND", []item{tAND, tEOF}},
	{"canNotThinkName", "(foo:value AND bar:value)", []item{
		item{itemLeftParen, 0, "("},
		item{itemAsk, 0, "foo:value"},
		item{itemAnd, 0, "AND"},
		item{itemAsk, 0, "bar:value"},
		item{itemRightParen, 0, ")"},
		tEOF,
	}},
}

// collect gathers the emitted items into a slice.
func collect(t *lexTest) (items []item) {
	l := lex(t.input)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func equal(i1, i2 []item, checkPos bool) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].typ != i2[k].typ {
			return false
		}
		if i1[k].val != i2[k].val {
			return false
		}
		if checkPos && i1[k].pos != i2[k].pos {
			return false
		}
	}
	return true
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		items := collect(&test)
		if !equal(items, test.items, false) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, items, test.items)
		}
	}
}
