package main

import (
	"fmt"
)

func sum(x int, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("x or y is less than 0")
	}
	return x + y, nil
}

func main() {
	result, error := sum(-1, 2)
	if error != nil {
		panic(error)
	}
	fmt.Println(result)
}
