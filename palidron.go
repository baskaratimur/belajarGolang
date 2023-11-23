package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Hapus spasi dan ubah ke huruf kecil
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Buat dua pointer
	front := 0
	back := len(s) - 1

	// Uji apakah string adalah palindrom
	for front < back {
		if s[front] != s[back] {
			return false // Jika tidak cocok, bukan palindrom
		}
		front++
		back--
	}
	return true // Jika semua pasangan karakter cocok, adalah palindrom
}

func main() {
	str := "madam1"
	if isPalindrome(str) {
		fmt.Println("String adalah palindrom.")
	} else {
		fmt.Println("String bukan palindrom.")
	}
}
