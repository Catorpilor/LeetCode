package suca

import (
	"fmt"
	"math"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func unsortedSubarray(nums []int) int {
	// return myWay(nums)
	// return useStack(nums)
	return fourPasses(nums)
}

// myWay time complexity O(N!)
func myWay(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	l, r := helper(nums)
	fmt.Printf("nums: %v, l: %d, r:%d \n", nums, l, r)
	if l != 0 && r != n-1 {
		return n
	} else if l == 0 && r == n-1 {
		return myWay(nums[l+1 : r])
	}
	if l == 0 {
		return myWay(nums[l+1:])
	}
	return myWay(nums[:r])
}

// helper returns nums's min & max position, time complexity O(N)
func helper(nums []int) (int, int) {
	n := len(nums)
	if n <= 1 {
		return 0, 0
	}
	l, r := 0, n-1
	minIdx, maxIdx := l, r
	for l <= r {
		if nums[l] < nums[minIdx] {
			minIdx = l
		}
		if nums[l] > nums[maxIdx] {
			maxIdx = l
		}
		if nums[r] < nums[minIdx] {
			minIdx = r
		}
		if nums[r] > nums[maxIdx] {
			maxIdx = r
		}
		l++
		r--
	}
	return minIdx, maxIdx
}

// selectionSort time complexity O(N^2), space complexity O(1)
func selectionSort(nums []int) int {
	n := len(nums)
	// just like selection sort, l, r are boundaries of the range needed to be sorted
	l, r := n, 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				// nums[i] and nums[j] are in the wrong position
				// update l, left most position
				l = utils.Min(l, i)
				// update r, right most position
				r = utils.Max(r, j)
			}
		}
	}
	if r < l {
		return 0
	}
	return r - l + 1
}

// useSort time complexity O(NlgN), space complexity O(N)
func useSort(nums []int) int {
	n := len(nums)
	local := make([]int, n)
	copy(local, nums)
	sort.Slice(local, func(i, j int) bool {
		return local[i] <= local[j]
	})
	l, r := n, 0
	for i := 0; i < n; i++ {
		if local[i] != nums[i] {
			l = utils.Min(l, i)
			r = utils.Max(r, i)
		}
	}
	if l > r {
		return 0
	}
	return r - l + 1
}

// useStack time complexity O(N), space complexity O(N)
func useStack(nums []int) int {
	n := len(nums)
	st := make([]int, 0, n)
	l, r := n, 0
	// traverse nums from left to right, if nums in ascending order, push index to st
	for i := 0; i < n; i++ {
		// if nums[i] < nums[i-1], means nums[i] at wrong pos, pop from stack to find the proper pos
		// and update l
		for len(st) > 0 && nums[st[len(st)-1]] > nums[i] {
			nst := len(st)
			l = utils.Min(l, st[nst-1])
			st = st[:nst-1]
		}
		st = append(st, i)
	}
	// clear
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i] {
			nst := len(st)
			r = utils.Max(r, st[nst-1])
			st = st[:nst-1]
		}
		st = append(st, i)
	}
	if l > r {
		return 0
	}
	return r - l + 1
}

// fourPasses time complexity O(N), space complexity O(1)
func fourPasses(nums []int) int {
	lMin, rMax := math.MaxInt32, math.MinInt32
	n := len(nums)
	// lMin represents the min when nums[i] < nums[i-1]
	// from left to right it should be ascending
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			// descending
			lMin = utils.Min(lMin, nums[i])
		}
	}
	// rMax represents the max when nums[i] > nums[i+1]
	// from right to left it should be descending,
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			// ascending
			rMax = utils.Max(rMax, nums[i])
		}
	}
	var l, r int
	for ; l < n; l++ {
		// find the proper postion for lMin
		if nums[l] > lMin {
			break
		}
	}
	for r = n - 1; r >= 0; r-- {
		// find the proper postion for rMax
		if nums[r] < rMax {
			break
		}
	}
	if l >= r {
		return 0
	}
	return r - l + 1
}
