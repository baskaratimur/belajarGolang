package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "baskara"
	names[1] = "baskara 1"
	names[2] = "baskara 2"

	fmt.Print(names[0])
	fmt.Print(names[1])
	fmt.Print(names[2])

	var angka = [3]int{
		90,
		95,
		85,
	}

	angkaa := [3]int{
		1, 2, 3,
	}
	fmt.Println(angkaa)
	fmt.Println(angkaa[1])
	fmt.Println(angka[0])

}
