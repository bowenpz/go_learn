package daily

var minVal int

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	minVal = root.Val
	return doFindSecondMinimumValue(root)
}

func doFindSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		if root.Val != minVal {
			return root.Val
		}
		return -1
	} else {
		left := doFindSecondMinimumValue(root.Left)
		right := doFindSecondMinimumValue(root.Right)
		if left == -1 && right == -1 {
			return -1
		} else if left == -1 {
			return right
		} else if right == -1 {
			return left
		} else {
			return min(left, right)
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
