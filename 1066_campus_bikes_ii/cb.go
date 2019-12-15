package cb

import (
	"fmt"
	"math"

	"github.com/catorpilor/LeetCode/utils"
)

func assignBikes(workers, bikes [][]int) int {
	nw, nb := len(workers), len(bikes)
	if nw == 0 || nb == 0 {
		return 0
	}
	res := math.MaxInt32
	visited := make([]bool, len(workers))
	dfs(&visited, bikes, workers, 0, 0, &res)
	return res
}

func dfs(visited *[]bool, bikes, workers [][]int, wid, distance int, res *int) {
	if wid >= len(workers) {
		*res = utils.Min(*res, distance)
		return
	}
	if distance > *res {
		return
	}
	for j := 0; j < len(bikes); j++ {
		if (*visited)[j] {
			continue
		}
		(*visited)[j] = true
		dfs(visited, bikes, workers, wid+1, distance+dis(bikes[j], workers[wid]), res)
		(*visited)[j] = false
	}
}

func dis(p1, p2 []int) int {
	return utils.Abs(p1[0]-p2[0]) + utils.Abs(p1[1]-p2[1])
}

func useDp(workers, bikes [][]int) int {
	nw, nb := len(workers), len(bikes)
	if nw == 0 || nb == 0 {
		return 0
	}
	res := math.MaxInt32
	dp := make([][]int, nw+1)
	// init dp
	for i := range dp {
		dp[i] = make([]int, 1<<nb)
		for j := 0; j < (1 << nb); j++ {
			dp[i][j] = math.MaxInt32 / 2
		}
	}
	for i := 1; i <= nw; i++ {
		for s := 1; s < (1 << nb); s++ {
			for j := 0; j < nb; j++ {
				if (s & (1 << j)) == 0 {
					continue
				}
				prev := s ^ (1 << j)
				dp[i][s] = utils.Min(dp[i][s], dp[i-1][prev]+dis(workers[i-1], bikes[j]))
				if i == nw {
					res = utils.Min(res, dp[i][s])
					fmt.Printf("all the worker are assigned, res:%d, dp[i][s]:%d\n", res, dp[i][s])
				}
			}
		}
	}
	return res
}
