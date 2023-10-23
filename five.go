package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}

	// Цикл for-range с массивом
	fmt.Println("Цикл for-range ")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)

	}
}
