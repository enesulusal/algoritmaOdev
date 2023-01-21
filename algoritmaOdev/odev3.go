package main

import "fmt"

func main() {
	var toplam int = 0
	for i := 0; i < 101; i++ {

		toplam += i
	}
	fmt.Println(toplam)

}
