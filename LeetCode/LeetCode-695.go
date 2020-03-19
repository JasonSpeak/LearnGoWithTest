package main

import "fmt"

var flag = [50][50]int{}

func maxAreaOfIsland(grid [][]int) int {
	flag = [50][50]int{}//每次调用flag前初始化flag，因为flag是全局变量，否则会受到前一个测试用例的影响
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if flag[i][j] == 1 {
				continue
			}
			if grid[i][j] == 1 {
				tmpArea := GetArea(grid, i, j) + 1
				if tmpArea > maxArea {
					maxArea = tmpArea
				}
			}
		}
	}
	return maxArea
}

func GetArea(grid [][]int, row, column int) int {
	area := 0
	width := len(grid[0])
	height := len(grid)
	flag[row][column] = 1
	if column-1 >= 0 {
		if grid[row][column-1] == 1 && flag[row][column-1] == 0 {
			area += GetArea(grid, row, column-1) + 1
		}
	}

	if column+1 < width {
		if grid[row][column+1] == 1 && flag[row][column+1] == 0 {
			area += GetArea(grid, row, column+1) + 1
		}
	}

	if row-1 >= 0 {
		if grid[row-1][column] == 1 && flag[row-1][column] == 0 {
			area += GetArea(grid, row-1, column) + 1
		}
	}

	if row+1 < height {
		if grid[row+1][column] == 1 && flag[row+1][column] == 0 {
			area += GetArea(grid, row+1, column) + 1
		}
	}
	return area
}

func main() {
	var grid = [][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {1, 1, 0, 1, 1}}
	fmt.Printf("%v", grid)
	area := maxAreaOfIsland(grid)
	fmt.Println(area)
}
