package main

import "fmt"

type hitung struct {
	i, j int
}

func main() {
	angka1 := hitung{42, 2703}
	angka2 := &angka1
	angka2.i = 21
	fmt.Println(angka1)
	angka3 := angka2.j / 37
	fmt.Println(angka3)
	fmt.Println(angka2)
}
