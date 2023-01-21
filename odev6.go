package main

import "fmt"

func main() {

	var average int = 0
	var total int = 0

	numbers := [8]int{89, 99, 86, 56, 90, 24, 15, 85}

	for i := 0; i < len(numbers); i++ {

		total += (numbers[i])
		average = (total / len(numbers))

	}
	fmt.Println(total, average)

}
