package main

import (
	"fmt"
)

func main() {
	sum := 0
	max := 260000
	for i := 1; i <= max; i++ {
		square := i * i
		if square%2 != 0 {
			sum += square
		}
	}
	fmt.Printf("%v\n", sum)
}
