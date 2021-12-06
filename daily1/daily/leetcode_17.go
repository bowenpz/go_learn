package daily

var charMap = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) (ans []string) {
	if len(digits) == 0 {
		return
	}

	var dfs func(i int, combination []byte)
	dfs = func(i int, combination []byte) {
		if i == len(digits) {
			ans = append(ans, string(combination))
		} else {
			for _, c := range charMap[digits[i]] {
				dfs(i+1, append(combination, byte(c)))
			}
		}
	}
	dfs(0, []byte{})
	return
}
