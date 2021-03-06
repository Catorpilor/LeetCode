package seq

import "sort"

func IsSub(s, t string) bool {
	m, n := len(s), len(t)
	if m == 0 {
		return true
	}
	if m > n {
		return false
	}
	idxs, idxt := 0, 0
	for idxt < n {
		if s[idxs] == t[idxt] {
			idxs += 1
			if idxs == m {
				return true
			}
		}
		idxt += 1
	}
	return false
}

func IsSubBS(s, t string) bool {
	m, n := len(s), len(t)
	if m == 0 {
		return true
	}
	if m > n {
		return false
	}
	dict := make(map[rune][]int)
	for i, c := range t {
		if dict[c] == nil {
			dict[c] = make([]int, 0, n)
		}
		dict[c] = append(dict[c], i)
	}
	prev := 0
	for _, c := range s {
		if dict[c] == nil {
			return false
		}
		tt := dict[c]
		j := sort.Search(len(tt), func(i int) bool {
			return tt[i] >= prev
		})
		if j == len(tt) {
			return false
		}
		prev = tt[j] + 1
	}
	return true
}
