package main

import "fmt"

type Address struct {
	gang, alamat, kota string
}

// pointer di functionecho“
// func getkota(adres *Address) {
// 	adres.kota = "kotapointer"
// }

// intinya kalau variable itu dia sebenarnya duplikat data
// bedanya sama pointer, dia ngubah datanya langsung, contoh adres 2 ngubah data ke adres1
// kalau variable dikasih pointer, maka adres1 bakal acuan ke adres 3
func main() {

	// variable struct
	var adres1 Address = Address{"depok", "almaatdepok", "kotadepok"}
	// adres1 := Address{"depok", "almaatdepok", "kotadepok"}

	// di pointer biar data Adressnya berubah, bukan diduplikat
	// adres2 := &adres1
	var adres2 *Address = &adres1

	adres2.gang = "bandung"

	// point 1 maka dia ngubah data adres1 jadi bandung, tapi adres2 dia buat data baru dan beda dari adres1
	// adres2 = &Address{"data2", "data22", "data222"}

	// point2 dia pointer di variable adres2, jadi siapapun yg make memory Adress bakal ke adres2 datanya
	*adres2 = Address{"data2", "data22", "data222"}

	fmt.Println(adres1)
	fmt.Println(adres2)

	// kalau ga dikasih pointer distruct sama pointer divariable function struct, maka munculnya cuman "" karna cmn duplikat , tapi kalau dikasih jadi keubah datanya kota pointer
	// alamat := &Address{"kota1", "kota2", ""}
	// getkota(alamat)
	// fmt.Println(alamat)

}
