package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	a := 10
	b := 20
	c := a + b
	// pake ausigment +c jadi ditambah lagi 30 + 20
	c += 20

	d := 4
	d++
	fmt.Println(c)
	fmt.Println(d)

	bas := "baskara"
	basbas := "baskara"
	var hasil = bas == basbas
	fmt.Println(hasil)
	// bisa juga < > !=

	nilai3 := 50
	nilai4 := 60

	fmt.Println(nilai3 > 40 && nilai4 > 70)
}
