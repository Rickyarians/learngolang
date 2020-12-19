package main

import "fmt"

func total(x int, y int) int {
	return x + y
}

func substract (x, y int) int {
	return x - y
}


func main() {
	fmt.Println(total(1, 2))
}