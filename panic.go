package main

import "fmt"

func exampleRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic telah terjadi:", r)
			// Lakukan sesuatu untuk menangani panic
		}
	}()
	panic("Ini adalah pesan error kritis!")
}

func main() {
	exampleRecover()
	fmt.Println("Program berlanjut setelah recover")
}
