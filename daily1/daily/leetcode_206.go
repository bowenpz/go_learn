package daily

// 递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.next == nil {
//		return head
//	}
//	newHead := reverseList(head.next)
//	head.next.next = head
//	head.next = nil
//	return newHead
//}

// 迭代
func reverseList(head *ListNode) *ListNode {
	pre, curr := head, head.Next
	pre.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
