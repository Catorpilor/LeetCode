package btzz

import "github.com/catorpilor/leetcode/utils"

func zigzagLevel(root *utils.TreeNode) [][]int {
	return bfsV2(root)
}

func bfs(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	res = append(res, []int{root.Val})
	leveld := make([]*utils.TreeNode, 0, 1000)
	leveld = append(leveld, root)
	zag := true
	for len(leveld) != 0 {
		next := make([]*utils.TreeNode, 0, 1000)
		for _, node := range leveld {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		lnodes := make([]int, len(next))
		for i, node := range next {
			lnodes[i] = node.Val
		}
		if zag {
			// swap vals
			for i, j := 0, len(next)-1; i < j; i, j = i+1, j-1 {
				lnodes[i], lnodes[j] = lnodes[j], lnodes[i]
			}
		}
		if len(lnodes) != 0 {
			res = append(res, lnodes)
		}
		leveld = next
		zag = !zag
	}
	return res
}

// bfsV2 time complexity is O(NlgN), space complexity is O(NlgN)
func bfsV2(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := make([]*utils.TreeNode, 0, 10000)
	q = append(q, root)
	zag := false
	for len(q) != 0 {
		l := len(q)
		next := make([]int, 0, 10000)
		for i := 0; i < l; i++ {
			if zag {
				next = append([]int{q[i].Val}, next...)
			} else {
				next = append(next, q[i].Val)
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[l:]
		zag = !zag
		res = append(res, next)
	}
	return res
}
