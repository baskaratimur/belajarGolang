package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("Baskara Timur"))
	fmt.Println(strings.ToUpper("Baskara Timur"))
	fmt.Println(strings.ToTitle("Baskara Timur"))
	fmt.Println(strings.Contains("Baskara Timur", "Baskara"))
	fmt.Println(strings.Contains("Baskara Timur", "Bara"))
	fmt.Println(strings.Split("Baskara Timur", " "))
	fmt.Println(strings.Trim("        Baskara Timur", " "))
	fmt.Println(strings.ReplaceAll("Baskara Timur", "Baskara", "Bara"))
}
