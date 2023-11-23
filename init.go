package main

import (
	"Belajar/database"
	"fmt"
)

func main() {
	db := database.Getdatabase()
	fmt.Println(db)

}
