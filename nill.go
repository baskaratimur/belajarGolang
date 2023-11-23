package main

import "fmt"

func getnill(namaa string) map[string]string {
	if namaa == "" {
		return nil
	} else {
		return map[string]string{
			"nama": namaa,
		}
	}

}
func main() {

	var orang map[string]string
	if orang == nil {
		fmt.Println("data kosong")

	} else {
		fmt.Println("data ada")

	}
}
