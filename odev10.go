package main

import (
	"fmt"
)

func main() {
	sayi := []int{2, 5, 26, 7, 5}
	adet := 0

	for i := 0; i < len(sayi); i++ {

		for j := 0; j < sayi[i]; j++ {
			if sayi[i] == sayi[j] {
				adet++
			}
		}

		fmt.Printf("%d tane %d elemanı vardır.\n", sayi, adet)
	}
}
