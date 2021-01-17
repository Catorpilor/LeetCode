package vowel

func countVowelString(n int) int {
	return useBruteForce(n)
}

// useBruteForce time complexity O(N^2), space complexity O(N)
func useBruteForce(n int) int {
	// n = 1 => 5
	// n = 2 => 15
	// a, e, i , o ,u, e, i, o ,u, i, o, u, o, u, u
	// n = 3
	// a -> a, e, i, o, u
	// e -> e, i ,o, u
	// i -> i, o, u
	// o -> o, u
	// u -> u
	// aa, ae, ai, ao, au, ee, ei, eo, eu, ii, io, iu, oo, ou, uu, ee, ei, eo,eu
	set := make(map[int]int, n)
	set[1] = 5
	set[2] = 15
	nd := []byte{'a', 'e', 'i', 'o', 'u', 'e', 'i', 'o', 'u', 'i', 'o', 'u', 'o', 'u', 'u'}
	factor := map[byte]int{'a': 5, 'e': 4, 'i': 3, 'o': 2, 'u': 1}
	for i := 3; i <= n; i++ {
		var local int
		for i := range nd {
			local += factor[nd[i]]
		}
		sb := make([]byte, 0, local)
		for i := range nd {
			switch nd[i] {
			case 'a':
				sb = append(sb, 'a', 'e', 'i', 'o', 'u')
			case 'e':
				sb = append(sb, 'e', 'i', 'o', 'u')
			case 'i':
				sb = append(sb, 'i', 'o', 'u')
			case 'o':
				sb = append(sb, 'o', 'u')
			case 'u':
				sb = append(sb, 'u')
			}
		}
		nd = sb
		set[i] = local
	}
	return set[n]
}
