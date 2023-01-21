package main

import "fmt"

func main() {

	sayilar := []int{52, 27, 81, 12, 99, 1}

	for i := 0; i < len(sayilar); i++ {
		for j := i + 1; j < len(sayilar); j++ {
			if sayilar[i] < sayilar[j] {
				sayilar[i], sayilar[j] = sayilar[j], sayilar[i]
			}
		}
	}

	fmt.Println(sayilar)
}
