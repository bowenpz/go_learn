package daily

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	reverse := func(bytes []byte, l, r int) {
		for l < r {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}
	}
	for i := 0; i < len(bytes); i += 2 * k {
		if i+k <= len(bytes) {
			reverse(bytes, i, i+k-1)
		} else {
			reverse(bytes, i, len(bytes)-1)
		}
	}
	return string(bytes)
}
