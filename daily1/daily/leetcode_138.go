package daily

func copyRandomList(head *Node) *Node {
	curr1 := head
	pre := new(Node)
	curr2 := pre

	nodeMap := make(map[*Node]*Node)

	for curr1 != nil {
		node := &Node{Val: curr1.Val}
		curr2.Next = node
		curr2 = node
		nodeMap[curr1] = node

		curr1 = curr1.Next
	}

	curr1 = head
	curr2 = pre.Next
	for curr2 != nil {
		if curr1.Random != nil {
			curr2.Random = nodeMap[curr1.Random]
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return pre.Next
}
