package main

import "fmt"

func main() {

	nama := "ilha1m"

	if nama == "ilham" {
		fmt.Println("ilham pinter")
	} else {
		fmt.Println("ilham bodoh")
	}
	namacas := "bara"
	panjangnama := len(namacas)
	switch {
	case panjangnama < 10:
		fmt.Println("Nama oke juga")
	case panjangnama > 10:
		fmt.Println("hai bere")
	default:
		fmt.Println("hai semuanya")

	}

	moons := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	// moonslen1 := moonslen
	// moonslen1 += 1
	for i := 0; i < len(moons); i++ {
		fmt.Println("data bulan ke", i+1, "=", moons[i])

	}

	// for i := 1; i <= 11; i++ {
	// 	fmt.Println("data bulan ke", i, "=", moons[i])

	// }

	// moonslen := len(moons)
	// // fmt.Print(moonslen)
	// for i := moonslen; i < 13; i++ {
	// 	fmt.Println(moons)
	// }

	for i := 1; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
