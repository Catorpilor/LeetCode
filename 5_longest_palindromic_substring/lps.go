package lps

// Lps returns the longest papindromic substring in s
func Lps(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, maxLen := 0, 1
	// for lenght 1
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// for length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			start, maxLen = i, 2
			dp[i][i+1] = true
		}
	}
	// for length >= 3
	for cl := 3; cl <= n; cl++ {
		for i := 0; i < n-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLen = cl
			}
		}
	}
	return s[start : start+maxLen]
}

// Lps2 returns the longest palindromic substring in s
func Lps2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	cap := 2*n + 1
	exs := make([]byte, cap)
	for i := range exs {
		if i%2 == 0 {
			exs[i] = '$'
		} else {
			exs[i] = s[i/2]
		}
	}

	dp := make([]int, cap)
	for i := range dp {
		dp[i] = 1
	}
	// prevcenter, leftbound, rightbound, maxlength of palindromic substring
	// nc stands for current center, fc stands for final palindromic substring center
	pc, lb, rb, max, nc, fc := 0, 0, 0, 1, 1, 0
	// i stands for current cent
	for nc < cap {
		cnt := maxPalinLength(exs, nc-dp[nc]/2-1, nc+dp[nc]/2+1)
		dp[nc] += cnt
		if dp[nc] > max {
			max = dp[nc]
			fc = nc
		}
		lb, rb = nc-dp[nc]/2, nc+dp[nc]/2
		if rb-lb <= 3 {
			nc += 1
		} else {
			// we have to find next center
			if rb == cap-1 {
				// if right bound reaches the end of string
				break
			}
			// set prev cent to current center
			pc = nc
			cmax := 0
			for j, k := nc+1, nc-1; j <= rb && k >= lb; j, k = j+1, k-1 {
				dp[j] = dp[k]
				// if k's left bound goes beyond current center's left bound
				// j can not be our next center
				if k-dp[k]/2 < lb {
					if k+dp[k]/2 > nc {
						dp[j] = rb - nc
					} else {
						dp[j] = 1
					}
				}
				if dp[j] > cmax && j+dp[j]/2 >= rb && k-dp[k]/2 >= lb {
					nc, cmax = j, dp[k]
				}
			}
			// if not found
			// nextcenter should be the right bound of current palindromic substring
			// or rb+1
			if pc == nc {
				nc = rb
			}
		}
	}
	start, end := fc-max/2, fc+max/2+1
	ret := make([]byte, 0, max/2)
	for i := start; i < end; i++ {
		if exs[i] != '$' {
			ret = append(ret, exs[i])
		}
	}
	return string(ret)
}

func maxPalinLength(b []byte, l, r int) int {
	count := 0
	for l >= 0 && r < len(b) && b[l] == b[r] {
		count, l, r = count+2, l-1, r+1
	}
	return count
}
