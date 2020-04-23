package sum

func subArraySum(nums []int, k int) int {
	return useBruteForce(nums, k)
}

// useBruteForce time complexity O(n^2), space complexity O(1)
func useBruteForce(nums []int, k int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	var ret int
	for i := range nums {
		// no need to set to nums[i]
		// cur := nums[i]
		// if cur == k {
		// 	ret++
		// }
		cur := 0
		for j := i; j < n; j++ {
			cur += nums[j]
			if cur == k {
				ret++
			}
		}
	}
	return ret
}

// useHashmMap time complexity O(N), space complexity O(N)
func useHashMap(nums []int, k int) int {
	var count, sum int
	// record[k]=v means how many sub-arrays sums equals to k
	// preSum for nums [-3,1,2,-3,5] is
	// preSum = [0,-3,-2,0,-3,2]
	// so to get SUM[1,3] which is [1,2,-3] we can use preSum to caculcate which is
	// preSum[4] - preSum[1] = -3 + 3 = 0
	// so if a subarray i to j sums equal to k , which means
	// preSum[j+1] - preSum[i] = k
	record := make(map[int]int)
	record[0] = 1
	for _, n := range nums {
		sum += n
		if v, ok := record[sum-k]; ok {
			count += v
		}
		record[sum]++
	}
	return count
}
