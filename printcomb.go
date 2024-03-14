package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a <= 2; a++ {
		for b := a + 1; b <= 3; b++ {
			for c := b + 1; c <= 4; c++ {
				z01.PrintRune('a')
				z01.PrintRune('b')
				z01.PrintRune('c')
				z01.PrintRune(',')
				z01.PrintRune(' ')
				z01.PrintRune(',')
			}
		}
	}
	z01.PrintRune('\n')
}
