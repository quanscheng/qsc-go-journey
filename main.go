package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("传参一", "传参二")
	fmt.Println(a, b)
}
