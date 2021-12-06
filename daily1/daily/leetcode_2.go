package daily

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	pre := new(ListNode)
	curr := pre
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		sum %= 10
		curr.Next = &ListNode{Val: sum}

		l1 = l1.Next
		l2 = l2.Next
		curr = curr.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		carry = sum / 10
		sum %= 10
		curr.Next = &ListNode{Val: sum}
		l1 = l1.Next
		curr = curr.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		carry = sum / 10
		sum %= 10
		curr.Next = &ListNode{Val: sum}
		l2 = l2.Next
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return pre.Next
}
