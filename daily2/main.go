package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	println(countAndSay(4))
}

func countAndSay(n int) string {
	str := "1"
	for n > 1 {
		n--
		preNum := str[0]
		preCnt := 1
		tmp := ""
		for i := range str[1:] {
			num := str[i+1]
			if preNum == num {
				preCnt++
			} else {
				tmp = fmt.Sprintf("%s%d%d", tmp, preCnt, preNum-'0')
				preCnt = 1
				preNum = num
			}
		}
		tmp = fmt.Sprintf("%s%d%d", tmp, preCnt, preNum-'0')
		str = tmp
	}
	return str
}

func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-2
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			ans[i-1] = "Fizz"
		} else if i%5 == 0 {
			ans[i-1] = "Buzz"
		} else {
			ans[i-1] = strconv.Itoa(i)
		}
	}
	return ans
}

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 11 {
		return []string{}
	}

	curr := int64(0)
	for _, c := range s[:10] {
		curr *= 10
		switch c {
		case 'A':
			curr += 1
		case 'C':
			curr += 2
		case 'G':
			curr += 3
		case 'T':
			curr += 4
		}
	}

	nums := make(map[int64]bool, 0)
	nums[curr] = true
	ans := make([]string, 0)
	ansMap := make(map[int64]bool, 0)

	for i, c := range s[10:] {
		curr = (curr % 1e9) * 10
		switch c {
		case 'A':
			curr += 1
		case 'C':
			curr += 2
		case 'G':
			curr += 3
		case 'T':
			curr += 4
		}
		if nums[curr] {
			if !ansMap[curr] {
				ansMap[curr] = true
				ans = append(ans, s[i+1:i+11])
			}
		} else {
			nums[curr] = true
		}
	}
	return ans
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	sum := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	cx1, cy1, cx2, cy2 := max(ax1, bx1), max(ay1, by1), min(ax2, bx2), min(ay2, by2)
	if cx1 < cx2 && cy1 < cy2 {
		sum -= (cx2 - cx1) * (cy2 - cy1)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func findMinMoves(machines []int) int {
	sum := 0
	for _, machine := range machines {
		sum += machine
	}
	if sum%len(machines) != 0 {
		return -1
	}

	avg := sum / len(machines)
	leftSum := 0
	ans := 0
	for _, num := range machines {
		num -= avg
		leftSum += num
		ans = max(ans, max(abs(leftSum), num))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(node *TreeNode, paths map[int64]int, sum int64) int
	dfs = func(node *TreeNode, paths map[int64]int, sum int64) (ans int) {
		if node == nil {
			return 0
		}
		sum += int64(node.Val)
		pathSum := sum - int64(targetSum)
		if cnt, ok := paths[pathSum]; ok {
			ans += cnt
		}
		paths[sum]++
		ans += dfs(node.Left, paths, sum)
		ans += dfs(node.Right, paths, sum)
		paths[sum]--
		return
	}
	return dfs(root, map[int64]int{0: 1}, 0)
}

func check1digit(ch byte) int {
	if ch == '*' {
		return 9
	}
	if ch == '0' {
		return 0
	}
	return 1
}

func check2digits(c0, c1 byte) int {
	if c0 == '*' && c1 == '*' {
		return 15
	}
	if c0 == '*' {
		if c1 <= '6' {
			return 2
		}
		return 1
	}
	if c1 == '*' {
		if c0 == '1' {
			return 9
		}
		if c0 == '2' {
			return 6
		}
		return 0
	}
	if c0 != '0' && (c0-'0')*10+(c1-'0') <= 26 {
		return 1
	}
	return 0
}

func numDecodings(s string) int {
	const mod int = 1e9 + 7
	a, b, c := 0, 1, 0
	for i := range s {
		c = b * check1digit(s[i]) % mod
		if i > 0 {
			c = (c + a*check2digits(s[i-1], s[i])) % mod
		}
		a, b = b, c
	}
	return c
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node.Child != nil {
			next := node.Next
			child := node.Child

			node.Child = nil
			node.Next = child
			child.Prev = node

			tail := dfs(child)
			tail.Next = next
			if next != nil {
				next.Prev = tail
			}

			node = tail
		}
		if node.Next != nil {
			return dfs(node.Next)
		} else {
			return node
		}
	}
	dfs(root)
	return root
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}
	pre := 1
	curr := 3
	for curr <= n && curr >= pre {
		if curr == n {
			return true
		}
		pre = curr
		curr *= 3
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) (ans []*ListNode) {
	node := head
	cnt := 0
	for node != nil {
		cnt++
		node = node.Next
	}

	node = head
	for node != nil {
		ans = append(ans, node)
		l := cnt / k
		if cnt%k != 0 {
			l++
		}
		cnt -= l
		k--
		for l > 1 && node.Next != nil {
			l--
			node = node.Next
		}
		tmp := node.Next
		node.Next = nil
		node = tmp
	}
	for k > 0 {
		k--
		ans = append(ans, nil)
	}
	return
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		rows := make(map[byte]bool, 0)
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				if !rows[board[i][j]] {
					rows[board[i][j]] = true
				} else {
					fmt.Printf("1 %d %d", i, j)
					return false
				}
			}
		}
	}
	for i := 0; i < len(board[0]); i++ {
		cols := make(map[byte]bool, 0)
		for j := 0; j < len(board); j++ {
			if board[j][i] != '.' {
				if !cols[board[j][i]] {
					cols[board[j][i]] = true
				} else {
					fmt.Printf("2 %d %d", i, j)
					return false
				}
			}
		}
	}
	latex := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		latex[i] = make(map[byte]bool, 0)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			idx := (i/3)*3 + (j / 3)
			if board[i][j] != '.' {
				if !latex[idx][board[i][j]] {
					latex[idx][board[i][j]] = true
				} else {
					fmt.Printf("3 %d %d", i, j)
					return false
				}
			}
		}
	}
	return true
}

type Trie struct {
	children [26]*Trie
	word     string
}

func (t *Trie) insert(word string) {
	trie := t
	for _, c := range word {
		if trie.children[c-'a'] == nil {
			trie.children[c-'a'] = &Trie{}
		}
		trie = trie.children[c-'a']
	}
	trie.word = word
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, word := range words {
		root.insert(word)
	}
	ansMap := make(map[string]bool, 0)

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int, t *Trie)
	dfs = func(i, j int, t *Trie) {
		if t != nil && board[i][j] != 0 {
			if len(t.word) > 0 {
				ansMap[t.word] = true
			}
			for _, dir := range dirs {
				nextI, nextJ := i+dir[0], j+dir[1]
				if nextI >= 0 && nextI < len(board) && nextJ >= 0 && nextJ < len(board[0]) && board[nextI][nextJ] != 0 {
					tmp := board[i][j]
					board[i][j] = 0
					dfs(nextI, nextJ, t.children[board[nextI][nextJ]-'a'])
					board[i][j] = tmp
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, root.children[board[i][j]-'a'])
		}
	}

	ans := make([]string, 0)
	for key := range ansMap {
		ans = append(ans, key)
	}
	return ans
}

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		} else {
			return dictionary[i] < dictionary[j]
		}
	})
	for _, dict := range dictionary {
		i, j := 0, 0
		for i < len(s) && j < len(dict) {
			for i < len(s) && s[i] != dict[j] {
				i++
			}
			if i < len(s) {
				i++
				j++
			}
		}
		if j == len(dict) {
			return dict
		}
	}
	return ""
}

func numberOfBoomerangs(points [][]int) (ans int) {
	for _, point := range points {
		x1, y1 := point[0], point[1]
		distanceMap := make(map[int]int, 0)
		for _, another := range points {
			x2, y2 := another[0], another[1]
			distance := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
			distanceMap[distance]++
		}
		for _, count := range distanceMap {
			ans += count * (count - 1)
		}
	}
	return
}

func chalkReplacer(chalk []int, k int) int {
	sum := chalk
	for i := 1; i < len(sum); i++ {
		sum[i] += sum[i-1]
	}
	k %= sum[len(sum)-1]
	for i, curr := range sum {
		if curr > k {
			return i
		}
	}
	return -1
}

func fullJustify(words []string, maxWidth int) (ans []string) {
	l, r, n := 0, 0, len(words)
	for r < n {
		cnt := len(words[l])
		for r+1 < n && cnt+len(words[r+1]) < maxWidth {
			r++
			cnt += len(words[r]) + 1
		}
		str := ""
		if r == n-1 {
			for l < n {
				if len(str) == 0 {
					str += words[l]
				} else {
					str += " "
					str += words[l]
				}
				l++
			}
			for len(str) < maxWidth {
				str += " "
			}
		} else if l == r {
			str += words[l]
			for len(str) < maxWidth {
				str += " "
			}
		} else {
			wordsSum := 0
			for i := l; i <= r; i++ {
				wordsSum += len(words[i])
			}
			spaceSum := maxWidth - wordsSum
			spaceCnt := r - l
			for l <= r {
				if len(str) == 0 {
					str += words[l]
				} else {
					insertSpace := spaceSum / spaceCnt
					if spaceSum%spaceCnt == 0 {
						for i := 0; i < insertSpace; i++ {
							str += " "
							spaceSum--
						}
					} else {
						for i := 0; i <= insertSpace; i++ {
							str += " "
							spaceSum--
						}
					}
					str += words[l]
					spaceCnt--
				}
				l++
			}
		}
		ans = append(ans, str)
		r++
		l = r
	}
	return
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	type item struct {
		p, c int
	}
	arr := make([]item, n)
	for i := 0; i < n; i++ {
		arr[i] = item{profits[i], capital[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].c < arr[j].c
	})

	h := &myHeap{}
	for i := 0; k > 0; k-- {
		for i < n && arr[i].c <= w {
			heap.Push(h, arr[i].p)
			i++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

type myHeap struct {
	sort.IntSlice
}

func (h *myHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *myHeap) Pop() interface{} {
	n := len(h.IntSlice)
	pop := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return pop
}

func (h myHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func balancedStringSplit(s string) int {
	cnt, sum := 0, 0
	bytes := []byte(s)
	for _, c := range bytes {
		if c == 'L' {
			cnt--
		} else {
			cnt++
		}
		if cnt == 0 {
			sum++
		}
	}
	return sum
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func compareVersion(version1 string, version2 string) int {
	bytes1, bytes2 := []byte(version1), []byte(version2)
	i, j := 0, 0

	var getNum = func(bytes []byte, i int) (num, newI int) {
		for i < len(bytes) {
			c := bytes[i]
			if c == '.' {
				break
			} else {
				num = num*10 + int(c-'0')
			}
			i++
		}
		return num, i
	}

	for i < len(bytes1) || j < len(bytes2) {
		num1, num2 := 0, 0
		num1, i = getNum(bytes1, i)
		num2, j = getNum(bytes2, j)

		if num1 == num2 {
			i++
			j++
		} else if num1 < num2 {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, booking := range bookings {
		ans[booking[0]-1] += booking[2]
		if booking[1] < n {
			ans[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

func numRescueBoats(people []int, limit int) (boats int) {
	sort.Ints(people)
	l, r := 0, len(people)-1
	for l < r {
		if people[l]+people[r] <= limit {
			l++
			r--
		} else {
			r--
		}
		boats++
	}
	if l == r {
		boats++
	}
	return
}
