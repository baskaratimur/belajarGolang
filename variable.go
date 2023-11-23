package main

import "fmt"

func main() {
	// using declaration jadi noktp tipe string, terus gue buat variable string
	type noKTp string
	var noktpbaskara noKTp = "123"
	fmt.Println(noktpbaskara)

	var name = "baskaratimur"
	fmt.Println(name)

	name = "barat"
	fmt.Println(name)

	name1 := "bara"
	fmt.Println(name1)

	var (
		multiple1 = "1"
		multiple2 = "2"
	)
	fmt.Println(multiple1)
	fmt.Println(multiple2)

	// const
	// const merupakan variable yg harus langsung ada datanya, sama ga bisa diubah,
	//  terus consta juga g masalah kalau ga dipake difile ini, g dianggap error kaya variable

	const (
		const1 = "bas"
		const2 = "bababas"
	)
	fmt.Println(const1)
}
