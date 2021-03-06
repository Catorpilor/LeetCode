package closest

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	return useHeap(points, k)
}

// An PointHeap is a max-heap of points .
type PointHeap [][]int

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// useHeap time complexity O(NlgK), space complexity O(K)
func useHeap(points [][]int, k int) [][]int {
	n := len(points)
	if k == n {
		return points
	}
	h := &PointHeap{}
	heap.Init(h)
	//  -10000 < points[0][0], points[0][1] < 10000
	// so the max distance is 10000*sqrt(2) = 15000
	for i := range points {
		heap.Push(h, points[i])
		for h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([][]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).([]int))
	}
	return ans
}

// useQuickSelect time complexity O(N), space complexity O(1)
func useQuickSelect(points [][]int, k int) [][]int {
	n := len(points)
	if n == k {
		return points
	}
	l, r := 0, n-1
	for l <= r {
		mid := helper(points, l, r)
		if mid == k {
			break
		} else if mid < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return points[:k]
}

func helper(points [][]int, l, r int) int {
	pivot := points[l]
	for l < r {
		for l < r && compare(points[r], pivot) >= 0 {
			r--
		}
		points[l] = points[r]
		for l < r && compare(points[l], pivot) <= 0 {
			l++
		}
		points[r] = points[l]
	}
	points[l] = pivot
	return l
}

func compare(p1, p2 []int) int {
	return p1[0]*p1[0] + p1[1]*p1[1] - p2[0]*p2[0] - p2[1]*p2[1]
}
