package main

import "fmt"

type HasName interface {
	getName() string
}

func sayhelo(hasName HasName) {
	fmt.Println("hallo", hasName.getName())
}

type Person struct {
	name string
}

func (person Person) getName() string {
	return person.name
}

func halobaskara(namabas string) {
	fmt.Println("hallo", namabas)
}

func main() {
	eko := Person{"baskara"}
	// eko.name = "baskara"
	sayhelo(eko)

	halobaskara(":baskara")

}
