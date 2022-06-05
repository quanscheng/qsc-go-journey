package main

import (
	"fmt"
)

// 多个返回值案例
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值案例
func split(sum int) (x, y int) {
	x = sum + 2
	y = sum - x
	return
}

func main() {
	a, b := swap("传参一", "传参二")

	fmt.Println(a, b)

	fmt.Println(split(6))
}
