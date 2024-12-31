package main

type Position struct {
	x int
	y int
}

func orangeRotten(grid [][]int) int {

	step := 0
	/*
	   | |
	  x|0|y
	   | |
	*/
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := make([]Position, 0)
	freshCount := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, Position{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	for len(queue) > 0 && freshCount > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, move := range moves {
				newX, newY := curr.x+move[0], curr.y+move[1]
				if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid) && grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					queue = append(queue, Position{newX, newY})
					freshCount--
				}
			}
		}
		step++
	}
	if freshCount > 0 {
		return -1
	}
	return step
}
