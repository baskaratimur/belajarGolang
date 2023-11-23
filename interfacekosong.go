package main

import "fmt"

func getdata(i int) interface{} {
	if i == 1 {
		return "baskara"
	} else if i == 2 {
		return 1
	} else {
		return true
	}
}

func main() {

	fmt.Println(getdata(1))

}
