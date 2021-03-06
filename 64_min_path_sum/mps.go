package mps

import (
	"fmt"

	"github.com/catorpilor/leetcode/utils"
)

func minPathSum(grid [][]int) int {
	// return useDp(grid)
	return useDpWithLessSpace(grid)
}

// useDp time complexity O(MN), space complexity O(MN)
func useDp(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if j == 0 {
			dp[0][j] = grid[0][0]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Min(dp[i-1][j]+grid[i-1][j], dp[i][j-1]+grid[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

// useDpWithLessSpace time complexity O(MN), space complexity O(N)
func useDpWithLessSpace(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			up, left := dp[j], dp[j-1]
			dp[j] = utils.Min(up+grid[i][j], left+grid[i][j])
		}
		fmt.Printf("row:%d, dp:%v\n", i, dp)
	}
	return dp[n-1]
}
