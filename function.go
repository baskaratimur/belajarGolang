package main

import "fmt"

func hello() {
	fmt.Println("helllowww")
}

// funct parameter
func alamat(Alamat string, Nohp int) {
	fmt.Println("dmana alamat?", Alamat, Nohp)
}

// funct return

func returnya(Nama string) string {
	return "Hello" + Nama
}

// func variadict
// di iterate pake range numbers
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// function didalem function
func cekfilter(Namanya string, spamfilterparam func(string) string) {
	fmt.Println("namanya ", spamfilter(Namanya))
}

func spamfilter(namanya string) string {
	if namanya == "anjing" {
		return "nama bahaya"
	} else {
		return namanya
	}
}

type filterSpam func(string) bool

// anonymous function lebih mudah
func cekfilter1(namanya1 string, paramfilter filterSpam) {
	if paramfilter(namanya1) {
		fmt.Println("baik")
	} else {
		fmt.Println("jahat")

	}

}

func main() {

	for i := 0; i < 5; i++ {
		hello()
	}

	// funct parameter
	alamat("depok", 7775838)

	// func return
	namanya := returnya("barabere")
	fmt.Println(namanya)

	// funct variadict
	jumlah := sumAll(10, 5, 5, 1, 2, 5)
	fmt.Println(jumlah)

	// function didalem function, jadinya function dibuat parameter
	cekfilter("anjing", spamfilter)

	varspamfilter := func(name string) bool {
		return name == "anjing"
	}

	cekfilter1("anjing", varspamfilter)
}
