package tree

import (
	"fmt"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func mctFromLeafValues(arr []int) int {
	return useDP(arr)
}

// useDP use top-down approach with time complexity O(N^3), space complexity O(N)
func useDP(arr []int) int {
	// res[i][j] means the min leaf values generated by nodes [i,j]
	set := make(map[string]int)
	return helper(arr, 0, len(arr)-1, set)
}

func helper(arr []int, l, r int, set map[string]int) int {
	key := fmt.Sprintf("%d-%d", l, r)
	if v, exists := set[key]; exists {
		return v
	}
	if l >= r {
		return 0
	}
	res := math.MaxInt32
	for i := l; i < r; i++ {
		root := max(arr[l:i+1]) * max(arr[i+1:r+1])
		res = utils.Min(res, root+helper(arr, l, i, set)+helper(arr, i+1, r, set))
	}
	set[key] = res
	return res
}

func max(arr []int) int {
	n := len(arr)
	if n < 1 {
		return 0
	}
	ans := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > ans {
			ans = arr[i]
		}
	}
	return ans
}

func minIdx(arr []int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	ans := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[ans] {
			ans = i
		}
	}
	return ans
}

// useGreedy time complexity O(N^2), space complexity O(1)
// we need to leave the max nodes close to the root to minimize the sum of non-leaf nodes.
// so the key point here is to use greedy to pick the smallest ones one at a time to build a
// subtree.
func useGreedy(arr []int) int {
	var ans int
	n := len(arr)
	for n > 1 {
		idx := minIdx(arr)
		// fmt.Printf("arr:%v, idx:%d\n", arr, idx)
		if idx > 0 && idx < n-1 {
			ans += arr[idx] * utils.Min(arr[idx-1], arr[idx+1])
		} else {
			if idx == 0 {
				ans += arr[idx] * arr[idx+1]
			} else {
				ans += arr[idx] * arr[idx-1]
			}
		}
		arr = append(arr[:idx], arr[idx+1:]...)
		n = len(arr)
	}
	return ans
}

// useStack time complexity O(N), space complexity O(N)
func useStack(arr []int) int {
	n := len(arr)
	// st is the monotonic decrasing stack, the top is the smallest.
	st := make([]int, 0, n)
	st = append(st, math.MaxInt32)
	var res int
	for _, num := range arr {
		nst := len(st)
		for nst > 0 && st[nst-1] <= num {
			mid := st[nst-1]
			st = st[:nst-1]
			nst--
			res += mid * utils.Min(num, st[nst-1])
		}
		st = append(st, num)
	}
	nst := len(st)
	for nst > 2 {
		top := st[nst-1]
		nst--
		st = st[:nst]
		res += top * st[nst-1]
	}
	return res
}
