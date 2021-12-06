package main

func main() {
	dict := Constructor()
	dict.AddWord("bad")
	println(dict.Search(".ad"))
}

func majorityElement(nums []int) (ans []int) {
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		numMap[num]++
	}
	size := len(nums) / 3
	for num, freq := range numMap {
		if freq > size {
			ans = append(ans, num)
		}
	}
	return
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0 && carry == 1; i-- {
		digits[i]++
		carry = digits[i] / 10
		digits[i] %= 10
	}
	if carry == 1 {
		newDigits := make([]int, 1, len(digits)+1)
		newDigits[0] = 1
		digits = append(newDigits, digits...)
	}
	return digits
}

func minMoves(nums []int) int {
	cnt := 0
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		cnt += num - min
	}
	return cnt
}

type Trie struct {
	Next   [26]*Trie
	Finish bool
}

type WordDictionary struct {
	Root *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &Trie{
			Next:   [26]*Trie{},
			Finish: false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.Root
	for _, c := range word {
		i := c - 'a'
		if node.Next[i] == nil {
			node.Next[i] = &Trie{
				Next:   [26]*Trie{},
				Finish: false,
			}
		}
		node = node.Next[i]
	}
	node.Finish = true
}

func (this *WordDictionary) Search(word string) bool {
	var doSearch func(node *Trie, word string) bool
	doSearch = func(node *Trie, word string) bool {
		for i, c := range word {
			if c == '.' {
				for _, nextNode := range node.Next {
					if nextNode != nil {
						if doSearch(nextNode, word[i+1:]) {
							return true
						}
					}
				}
				return false
			} else {
				node = node.Next[c-'a']
				if node == nil {
					return false
				}
			}
		}
		return node.Finish
	}
	return doSearch(this.Root, word)
}
