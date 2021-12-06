package daily

func calculate(s string) int {
	x, y := 1, 0
	for _, c := range s {
		if c == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}
