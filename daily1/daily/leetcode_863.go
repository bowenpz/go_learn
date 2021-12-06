package daily

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 记录所有节点的父节点
	parents := make(map[int]*TreeNode)
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)

	// 找距离是 k 的节点
	ans := make([]int, 0)
	var findKNode func(*TreeNode, *TreeNode, int)
	findKNode = func(node *TreeNode, from *TreeNode, k int) {
		if k == 0 {
			ans = append(ans, node.Val)
			return
		} else {
			if parent := parents[node.Val]; parent != nil && parent != from {
				findKNode(parent, node, k-1)
			}
			if left := node.Left; left != nil && left != from {
				findKNode(left, node, k-1)
			}
			if right := node.Right; right != nil && right != from {
				findKNode(right, node, k-1)
			}
		}
	}
	findKNode(target, nil, k)

	return ans
}
