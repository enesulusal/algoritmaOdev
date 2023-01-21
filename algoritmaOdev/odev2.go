package main

import "fmt"

func main() {
	var sayi int
	fmt.Print("sayi girin: ")
	fmt.Scan(&sayi)

	if sayi%2 == 0 {
		fmt.Println(sayi, "çift sayi")
	} else {
		fmt.Println(sayi, "tek sayi")
	}
}
