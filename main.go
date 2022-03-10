package main

import (
	"flag"
	"fmt"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/jsontogo/jsontogo"
)

var (
	typeName = flag.String("typename", "foo", "the typename")
)

func main() {
	flag.Parse()
	for _, a := range flag.Args() {
		s, err := jsontogo.JSONToGo(a, *typeName)
		check.Err(err)
		fmt.Println(s)
	}
}
