package main

import (
	"fmt"
)

func main() {
	p := 5000.0
	t := 1.0
	r := 10.0
	s := simpleinterest.si(p, t, r)
	fmt.Printf("Simple interest calculation = ", s)
}
