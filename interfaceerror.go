package main

import (
	"errors"
	"fmt"
)

func Pembagian(i int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("data ga bisa dibagi kosong")
	} else {
		return i / pembagi, nil
	}
}

func main() {

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("hasilnya", hasil)

	} else {
		fmt.Println("Error", err.Error())
	}
}
