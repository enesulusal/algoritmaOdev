package main

import (
	"fmt"
)

func main() {
	sayi := []int{2, 5, 26, 7, 5}
	for i := 0; i < len(sayi); i++ {
		adet := 0
		for j := 0; j < len(sayi); j++ {
			if sayi[i] == sayi[j] {
				//fmt.Println(sayi[i])
				adet++

			}
		}
		fmt.Printf("%d adet %d elemanı vardır.\n", adet, sayi[i])

	}
}
