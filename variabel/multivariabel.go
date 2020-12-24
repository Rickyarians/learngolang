package main

import "fmt"

func main() {

	// go mendukung deklarasi multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	// lenbih ringkas
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// lebih ringkas lagi
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(seventh, eight, ninth)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(one, isFriday, twoPointTwo, say)

}
