package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	integer, err := strconv.ParseInt("1000", 10, 64)
	if err == nil {
		fmt.Println(integer)
	} else {
		fmt.Println(err.Error())
	}

	inttotrue := strconv.FormatInt(10000, 10)
	fmt.Println(inttotrue)

	fmt.Println(math.Round(104.7))
	fmt.Println(math.Floor(104.3))
	fmt.Println(math.Ceil(104.3))

}
