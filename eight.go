package main

import "fmt"

func main() {
	sample := "слово"

	fmt.Println("Both Index and Value")
	for i, letter := range sample {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}

}
