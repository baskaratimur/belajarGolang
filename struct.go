package main

import "fmt"

type Customer struct {
	name, address string
	umur          int
}

// struct method
func (customer Customer) sayHi(name string) {
	// name diambil dari param name di sayhi
	fmt.Println("say hallo", name, "dimana alamat", customer.address)
}

func main() {
	// struct biasa
	// eko := Customer{}
	// eko.name = "baskara"
	// eko.address = "depok"
	// eko.umur = 15
	// fmt.println(eko)

	// atau
	// eko := Customer{"baskara", "depopk", 15}
	eko := Customer{"bara", "depok", 15}
	eko.sayHi("baskara")
}
