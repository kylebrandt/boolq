package boolq

import (
	"fmt"
	"testing"
)

type simpleAsker struct{}

func (s simpleAsker) Ask(f string) (bool, error) {
	switch f {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("not a valid simpleAsker ask: %v", f)
	}
}

const (
	hasError = true
	noError  = false
)

type simpleAskerTest struct {
	input     string
	result    bool
	shouldErr bool
}

var tests = []simpleAskerTest{
	{"", false, hasError},
	{"true", true, noError},
	{"!true", false, noError},
	{"!false", true, noError},
	{"!false AND true", true, noError},
}

func TestBoolQ(t *testing.T) {
	var sa simpleAsker
	for _, test := range tests {
		q, err := AskExpr(test.input, sa)
		if err == nil == test.shouldErr {
			t.Errorf("unexpected error in expr %v: %v", test.input, err)
		}
		if q != test.result {
			t.Errorf("unexpted result for expr %v: got %v expected %v", test.input, q, test.result)
		}
	}
}
