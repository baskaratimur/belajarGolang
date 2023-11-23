package main

import (
	"fmt"
)

func ApakahPalidron(teks string) bool {
	// teks = strings.ToLower(teks) // Ubah teks menjadi huruf kecil
	depan := 0
	belakang := len(teks) - 1

	for depan < belakang {
		if teks[depan] != teks[belakang] {
			return false
		}
		depan++
		belakang--
	}
	return true

}

func main() {
	kata := "ibi"
	if ApakahPalidron(kata) {
		fmt.Println("palidron")
	} else {
		fmt.Println("bukan Palidron")
	}
}
