package main

import "fmt"

func main() {
	var vize, final float32
	fmt.Print("Vize notunuzu giriniz: ")

	fmt.Scan(&vize)

	fmt.Print("Final notunuzu giriniz: ")
	fmt.Scan(&final)
	ortalama := (vize * 0.3) + (final * 0.7)
	fmt.Println("Ortalama:", ortalama)
}
