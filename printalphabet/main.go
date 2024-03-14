package main

import "github.com/01-edu/z01"

func main() {
	for b := 97; b <= 122; b++ {
		z01.PrintRune(rune(b))
	}
	z01.PrintRune('\n')
}
