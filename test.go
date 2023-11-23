package main

import "fmt"

type filterSpam12 func(string) bool

// anonymous function ceritanya
func cekfilter12(nama string, filterspamparam filterSpam12) {
	if filterspamparam(nama) {
		fmt.Println("baik")
	} else {
		fmt.Println("jahat")

	}

}

func main() {
	varfilterSpam3 := func(name string) bool {
		return name == "anjing"
	}
	cekfilter12("anjing", varfilterSpam3)
}

// defer menjalankan function walau gagal
// panic menghentikan program, jadi yg dibawah kaga dirunning
// recover buat tetap menjalankan program, nnti errornya ditangkap, contoh
// defer endapp(){}
// // nah nnti difunction defer
// messageerror := recover()
// fmt.Println("aplikasi error", messageerror)

// runapp()
// if error {
// 	panic ("aplikasi error")
// }
