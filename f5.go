package main

import "fmt"

func main() {
	add(1, 2, 3)
	add(1, 2, 3, 4)
	add(5, 6, 7, 2, 3)
}

func add(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}
