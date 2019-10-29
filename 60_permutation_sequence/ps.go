package ps

import (
	"strconv"
	"strings"
)

func getPermutation(n, k int) string {
	if n < 1 {
		return ""
	}
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	res := make([]string, 0, k)
	genRes(&res, nums)

	// var sb strings.Builder
	// for i := 0; i < n; i++ {
	// 	sb.WriteString(strconv.Itoa(i))
	// }
	// src := sb.String()
	// res := make([]string, 0, k)
	// res = append(res, src)
	// permute(&res, "", 0, src)
	return res[k-1]
}

// just call nextPermutation k times
// src are monotonically increasing "123456...n"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] > nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] < nums[i] {
			j--
		}
	}
}

// func nextPermutation(src string) string{
// 	n := len(src)
// 	i := n - 1
// 	for i >=0 && src[i-1] > src[i] {
// 		i--
// 	}
// 	if i > = 0 {
// 		j :=
// 	}
// }

func genRes(res *[]string, nums []int) {
	var st strings.Builder
	for i := range nums {
		st.WriteString(strconv.Itoa(nums[i]))
	}
	*res = append(*res, st.String())
}

// func permute(res *[]string, prefix string, pos int, s string) {
// 	if pos == len(s) {
// 		*res = append(*res, prefix)
// 		return
// 	}
// 	for i := pos; i < len(s); i++ {

// 	}
// }
