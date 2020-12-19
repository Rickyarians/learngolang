// exported public akses, ditandai dengan huruf besar
// unexported private akses, ditandai huruf kecil


package main

import (
	"fmt"
	"github.com/rickyarians/belajargolang/aritmethic"
)



func substract (x, y int) int {
	return x - y
}


func main() {
	fmt.Println(total(1, 2))
}