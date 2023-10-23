package main

import "fmt"

func main() {
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	fmt.Println(" Key and Value")
	for k, v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v)
	}

}
