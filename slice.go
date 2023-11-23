package main

import "fmt"

func main() {
	array1 := [...]string{
		"bas0",
		"bas1",
		"bas2",
		"bas3",
		"bas4",
		"bas5",
		"bas6",
		"bas7",
		"bas8",
		"bask9",
	}
	slice := array1[3:6]
	slice[0] = "4444"
	fmt.Println(slice)

	// fungsi buat nambah data baru diarray
	sliceappend := array1[8:]
	sliceappend1 := append(slice, "bask10")
	sliceappend1[1] = "999"
	fmt.Println(sliceappend1)
	fmt.Println(sliceappend)

	// cap = ngitung kapasitas dari si slice
	// cap1 := cap(slice)
	// fmt.Println(slice1, cap(slice1))

	// make itu untuk membuat slice
	slicemake := make([]string, 4)
	slicemake[0] = "memek"
	slicemake[1] = "memek1"
	slicemake[2] = "memek2"
	slicemake[3] = "memek3"

	fmt.Println(slicemake)
}
