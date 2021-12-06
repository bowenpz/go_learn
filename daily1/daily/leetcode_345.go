package daily

var vowelMap = map[byte]bool{
	'a': true,
	'e': true,
	'l': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	reverse := []byte(s)
	for l, r := 0, len(s)-1; l < r; {
		for l < r && !vowelMap[reverse[l]] {
			l++
		}
		for l < r && !vowelMap[reverse[r]] {
			r--
		}
		reverse[l], reverse[r] = reverse[r], reverse[l]
		l++
		r--
	}
	return string(reverse)
}
