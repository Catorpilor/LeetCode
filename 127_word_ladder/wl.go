package wl

import "fmt"

func ladderLength(beg, end string, lists []string) int {
	return bfs(beg, end, lists)
}

// bfs time complexity O(N^2*M) where M is the length of lists[i]
// space complexity is O(N)
func bfs(beg, end string, lists []string) int {
	n := len(lists)
	if n < 1 {
		return 0
	}
	hset := make(map[string][]string, len(lists))
	helper(hset, lists)
	// withPattern(hset, lists, len(beg))
	if _, exists := hset[end]; !exists {
		return 0
	}
	// level 1
	level := 1
	visited := make(map[string]int, len(lists))
	leveld := make([]string, 0, len(lists))
	for i := range lists {
		if diffOne(beg, lists[i]) {
			leveld = append(leveld, lists[i])
			visited[lists[i]] = level + 1
		}
		if beg == lists[i] {
			visited[beg] = level
		}
	}
	level++
	fmt.Printf("beg's diffOne resule: %v\n", leveld)
	for len(leveld) != 0 {
		next := make([]string, 0, len(lists))
		for i := range leveld {
			if leveld[i] == end {
				return level
			}
			for _, str := range hset[leveld[i]] {
				if _, exists := visited[str]; exists {
					continue
				}
				next = append(next, str)
				visited[str] = level + 1
			}
		}
		leveld = next
		level++
		fmt.Printf("next level:%d, with vals: %v\n", level, leveld)
	}
	return 0
}

func withPattern(hset map[string][]string, lists []string, l int) {
	for _, str := range lists {
		for i := 0; i < l; i++ {
			key := str[:i] + "*" + str[i+1:]
			hset[key] = append(hset[key], str)
		}
	}
}

func helper(hset map[string][]string, lists []string) {
	for i := range lists {
		hset[lists[i]] = make([]string, 0, len(lists))
		for j := range lists {
			if i == j {
				continue
			}
			if diffOne(lists[i], lists[j]) {
				hset[lists[i]] = append(hset[lists[i]], lists[j])
			}
		}
	}
}

func diffOne(a, b string) bool {
	diffs := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs == 1
}

// bfsWithPattern time complexity  O(M*N) M is len(lists), N is len(beg)
// space complexity O(M)
func bfsWithPattern(beg, end string, lists []string) int {
	hset := make(map[string][]string, len(lists))
	withPattern(hset, lists, len(beg))
	visited := make(map[string]bool, len(lists))
	leveld := make([]string, 0, len(lists))
	for i := 0; i < len(beg); i++ {
		key := beg[:i] + "*" + beg[i+1:]
		leveld = append(leveld, hset[key]...)
	}
	visited[beg] = true
	level := 2
	for len(leveld) != 0 {
		next := make([]string, 0, len(lists))
		for _, str := range leveld {
			if visited[str] {
				continue
			}
			visited[str] = true
			if str == end {
				return level
			}
			for i := 0; i < len(beg); i++ {
				key := str[:i] + "*" + str[i+1:]
				next = append(next, hset[key]...)
			}
		}
		level++
		leveld = next
	}
	return 0
}
