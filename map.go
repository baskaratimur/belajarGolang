package main

import "fmt"

func main() {
	data1 := map[string]string{
		"nama":   "baskara",
		"alamat": "depok",
	}

	data1["nama"] = "bara"
	fmt.Println(data1)
	fmt.Println(data1["nama"])

	mapmake := make(map[string]string)
	mapmake["buku"] = "naruto"
	mapmake["gender"] = "laki"

	fmt.Println(mapmake)
	delete(mapmake, "gender")
	fmt.Println(mapmake)

}
