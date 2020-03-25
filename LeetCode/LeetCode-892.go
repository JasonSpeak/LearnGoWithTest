package main

// 在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。
// 每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。
// 请你返回最终形体的表面积。
// 示例 1：
// 输入：[[2]]
// 输出：10
// 示例 2：
// 输入：[[1,2],[3,4]]
// 输出：34
// 示例 3：
// 输入：[[1,0],[0,2]]
// 输出：16
// 示例 4：
// 输入：[[1,1,1],[1,0,1],[1,1,1]]
// 输出：32
// 示例 5：
// 输入：[[2,2,2],[2,1,2],[2,2,2]]
// 输出：46
// 提示：
// 1 <= N <= 50
// 0 <= grid[i][j] <= 50

func surfaceArea(grid [][]int) int {
	h := len(grid)
	if h == 0 {
		return 0
	}

	w := len(grid[0])
	if w == 0 {
		return 0
	}

	m := func(x, y int) int {
		min := x
		if y < x {
			min = y
		}
		return min
	}
	ret := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] > 0 {
				ret += 2 //top and bottom

				cur := grid[i][j]
				ret += 4 * cur //common

				// - up
				if i != 0 {
					cp := grid[i-1][j]
					min := m(cur, cp)
					ret -= min * 2
				}

				// - left
				if j != 0 {
					cp := grid[i][j-1]
					min := m(cur, cp)
					ret -= min * 2
				}
			}
		}
	}

	return ret
}
