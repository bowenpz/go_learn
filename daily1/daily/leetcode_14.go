package daily

func longestCommonPrefix(strs []string) string {
	ans := make([]byte, 0)
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, str := range strs {
			if len(str) == i || str[i] != c {
				return string(ans)
			}
		}
		ans = append(ans, c)
	}
	return string(ans)
}
