package number

func numsSameDiff(n, k int) []int {
	return useBruteForce(n, k)
}

// useBruteForce time complexity O(1), space compleixty O(N)?
func useBruteForce(n, k int) []int {
	base := make([]int, 10)
	for i := range base {
		base[i] = i
	}
	for p := 2; p <= n; p++ {
		// fmt.Printf("base:%v, pre:%d\n", base, pre)
		tmp := []int{}
		for i := range base {
			orig := base[i] % 10
			// fmt.Printf("orig:%d, base[i]:%d\n", orig, base[i])
			if base[i] > 0 && orig >= k {
				tmp = append(tmp, base[i]*10+orig-k)
			}
			// only add when k!=0 remove duplicates
			if base[i] > 0 && orig+k <= 9 && k != 0 {
				tmp = append(tmp, base[i]*10+orig+k)
			}
		}
		base = tmp
	}
	return base
}
