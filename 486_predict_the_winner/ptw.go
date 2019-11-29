package ptw

import (
	"math"

	"github.com/catorpilor/LeetCode/utils"
)

func predictTheWinner(nums []int) bool {
	// return helper(nums, 0, len(nums)-1) >= 0
	n := len(nums)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = math.MinInt32
		}
	}
	return withMemo(nums, 0, n-1, cache) >= 0
}

// helper implements maxmin algorithm
// each turn player has two choices nums[s] or nums[e]
// player 1 wins ==> sum of #1 picks >= sum of #2 picks
// aka. sum of #1 picks - sum of #2 picks >= 0
// so when player 1 picks we add it to the total sum, then it's player 2's turn we just subtract it form the total sum
// eg. player 1 picks nums[s] , player 2's picks is helper(nums, s+1,e)
// time complexity is O(2^n) each time we have two choices.
// space complexity is O(n), the recursion tree can only be n length deep.
func helper(nums []int, s, e int) int {
	if s == e {
		return nums[e]
	}
	return utils.Max(nums[s]-helper(nums, s+1, e), nums[e]-helper(nums, s, e-1))
}

// withMemo implements maxmin algorithm
// time complexity O(n^2)
// space complexity O(n^2)
func withMemo(nums []int, s, e int, mem [][]int) int {
	if mem[s][e] == math.MinInt32 {
		if s == e {
			mem[s][e] = nums[e]
		} else {
			mem[s][e] = utils.Max(nums[s]-withMemo(nums, s+1, e, mem), nums[e]-withMemo(nums, s, e-1, mem))
		}
	}
	return mem[s][e]
}

func dynamic(nums []int) bool {
	n := len(nums)
	if n <= 1 || n&1 == 0 {
		return true
	}
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i] = nums[i]
			} else {
				dp[j] = utils.Max(nums[i]-dp[j], nums[j]-dp[j-1])
			}
		}
	}
	return dp[n-1] >= 0
}
