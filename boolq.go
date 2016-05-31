// Package boolq lets you build generic query expressions.

package boolq

import (
	"fmt"

	"github.com/kylebrandt/boolq/parse"
)

type Asker interface {
	Ask(string) (bool, error)
}

func AskExpr(expr string, asker Asker) (bool, error) {
	q, err := parse.Parse(expr)
	if err != nil {
		return false, err
	}
	return walk(q.Root, asker)
}

func walk(node parse.Node, asker Asker) (bool, error) {
	switch node := node.(type) {
	case *parse.AskNode:
		return asker.Ask(node.Text)
	case *parse.BinaryNode:
		return walkBinary(node, asker)
	case *parse.UnaryNode:
		return walkUnary(node, asker)
	default:
		panic(fmt.Errorf("can't walk this type", node))
	}
	return true, nil
}

func walkBinary(node *parse.BinaryNode, asker Asker) (bool, error) {
	l, err := walk(node.Args[0], asker)
	if err != nil {
		return false, err
	}
	r, err := walk(node.Args[1], asker)
	if err != nil {
		return false, err
	}
	if node.OpStr == "AND" {
		return l && r, nil
	}
	if node.OpStr == "OR" {
		return l || r, nil
	}
	return false, fmt.Errorf("Unrecognized operator: %v", node.OpStr)
}

func walkUnary(node *parse.UnaryNode, asker Asker) (bool, error) {
	r, err := walk(node.Arg, asker)
	if err != nil {
		return false, err
	}
	if node.OpStr == "!" {
		return !r, nil
	}
	return false, fmt.Errorf("unknown unary operator: %v", node.OpStr)
}
