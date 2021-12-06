package daily

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	usedA, usedB := false, false
	currA, currB := headA, headB
	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
		if currA == nil && !usedA {
			usedA = true
			currA = headB
		}
		if currB == nil && !usedB {
			usedB = true
			currB = headA
		}
	}
	return nil
}
