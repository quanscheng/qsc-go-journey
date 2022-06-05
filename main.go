package main

import (
	"fmt"
	"qsc-go-journey/user"
)

func main() { // `entry` : main function with main package
	s := user.Sum(4, 5) // user.Hello().var! 回车 就可以快捷生成变量
	fmt.Printf("s: %v\n", s)
}
