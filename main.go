package main

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Sum(3, 4))
	fmt.Println(sum3(1, 2, 3))
}

func sum3(a int, b int, c int) int {
	return a + b + c
}
