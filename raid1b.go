package student

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						z01.PrintRune('/')
					} else if j == x {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						z01.PrintRune('\\')
					} else if j == x {
						z01.PrintRune('/')
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			} else {
				for j := 1; j <= x; j++ {
					if j == x || j == 1 {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
	}
}
