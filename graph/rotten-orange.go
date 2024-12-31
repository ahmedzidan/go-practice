package main

import "fmt"

func getRottenOrange(grid [][]int) bool {
	fresh := 0
	queue := make([][]int, 0)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i, row := range grid {
		for j, val := range row {
			if val == 2 {
				queue = append(queue, []int{i, j})
			} else if val == 1 {
				fresh++
			}
		}
	}

	for len(queue) > 0 && fresh != 0 {
		current := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			x, y := current[0]+dir[0], current[1]+dir[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid) && grid[x][y] == 2 {
				grid[x][y] = 2
				fresh--
				queue = append(queue, []int{x, y})
			}
		}
	}
	if fresh > 0 {
		return false
	}

	return true
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 2, 2},
		{2, 1, 1},
	}
	fmt.Println(getRottenOrange(grid))
}
