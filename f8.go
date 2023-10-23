package main

import "fmt"

func totoshka(x int, y int) int {

	return x + y
}
func rogalik(x int, y int) int {

	return x * y
}
func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}
func main() {

	action(100, 25, totoshka)
	action(50, 60, rogalik)
}
