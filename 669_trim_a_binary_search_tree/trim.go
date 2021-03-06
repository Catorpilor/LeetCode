package trim

import "github.com/catorpilor/leetcode/utils"

func trimBST(root *utils.TreeNode, low, high int) *utils.TreeNode {
	return helper(root, low, high)
}

// helper time complexity O(N), space complexity O(log(N))
func helper(root *utils.TreeNode, low, high int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return helper(root.Right, low, high)
	}
	if root.Val > high {
		return helper(root.Left, low, high)
	}
	root.Left = helper(root.Left, low, high)
	root.Right = helper(root.Right, low, high)
	return root
}
