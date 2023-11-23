package main

import "fmt"

type Poinmethod struct {
	name string
}

func (poinmethod *Poinmethod) callinguser() {
	poinmethod.name = "Namanya" + poinmethod.name
}

func main() {
	var baskara = Poinmethod{"baskara"}
	baskara.callinguser()
	fmt.Println(baskara.name)
}
