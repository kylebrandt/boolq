# boolq
build simple bool expressions

# Example:

```
package main

import (
	"fmt"
	"log"

	"github.com/kylebrandt/boolq"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	f := foo{}
	ask := "(true AND true) AND !false"
	q, err := boolq.AskExpr(ask, f)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(q)
}

type foo struct{}

func (f foo) Ask(ask string) (bool, error) {
	switch ask {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, fmt.Errorf("couldn't parse ask arg")
}
```