package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	fmt.Println(math.Sqrt(28))

	s := time.Now()
	t := time.Now().Second()
	fmt.Println("Текущее время: ", s, "Секунды:", t)

}
