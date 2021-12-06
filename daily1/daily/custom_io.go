package daily

func valid(password string) bool {
	bytes := []byte(password)

	if !isLetter(bytes[0]) {
		return false
	}

	hasNum := false
	for _, c := range bytes {
		if isNum(c) {
			hasNum = true
		} else if !isLetter(c) {
			return false
		}
	}
	return hasNum
}

func isLetter(c byte) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}
