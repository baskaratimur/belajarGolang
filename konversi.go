package main

import "fmt"

func main() {
	var nilai32 int32 = 3278
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	// ambil string
	var baskara = "baskara timur"
	var e = baskara[9]
	var eEtring string = string(e)

	fmt.Println(eEtring)
}
