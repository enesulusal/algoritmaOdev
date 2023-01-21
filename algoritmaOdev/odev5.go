package main

import "fmt"

func main() {
	sayilar := []int{12, 15, 79, 34, 97, 7}
	min := sayilar[0]

	for i := 1; i < len(sayilar); i++ {
		if sayilar[i] < min {
			min = sayilar[i]
		}
	}
	fmt.Println("en küçük: ", min)
}
