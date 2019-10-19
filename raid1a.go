package student

import "github.com/01-edu/z01"

func Raid1a(x,y int) {
	for i := 1; i <= y; i++ {
		if i == y || i == 1 {
			for j := 1; j <= x; j++ {
				if j == x || j == 1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		} else {
			for j := 1; j <= x; j++ {
				if j == x || j == 1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}