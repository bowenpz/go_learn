package daily

func checkRecord(s string) bool {
	aCnt, lCnt := 0, 0
	for _, c := range s {
		if c == 'A' {
			aCnt++
			lCnt = 0
		} else if c == 'L' {
			lCnt++
		} else {
			lCnt = 0
		}

		if aCnt == 2 || lCnt == 3 {
			return false
		}
	}
	return true
}
