package main

import "fmt"

// convert tipe data

func returndata() interface{} {
	return "ok"
}
func main() {
	result := returndata()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")

	}
}
