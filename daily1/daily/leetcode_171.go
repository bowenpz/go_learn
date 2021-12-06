package daily

func titleToNumber(columnTitle string) int {
	num := 0
	for _, c := range columnTitle{
		num *= 26
		num += int(c - 'A' + 1)
	}
	return num
}
