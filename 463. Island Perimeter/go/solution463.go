package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(grid))
}

type cell struct {
	x int
	y int
}

type queue []cell

func (q *queue) push(c cell) {
	*q = append(*q, c)
}

func (q *queue) pop() cell {
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}

func islandPerimeter(grid [][]int) int {
	rows := len(grid)    // number of rows
	cols := len(grid[0]) // number of cols

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				q := &queue{}
				q.push(cell{x: c, y: r})
				visited[r][c] = true
				perimeter := 0

				directions := [][]int{
					{0, 1},
					{1, 0},
					{0, -1},
					{-1, 0},
				}

				for len(*q) != 0 {
					curr := q.pop()
					for _, dir := range directions {
						dx, dy := dir[0], dir[1]
						nx, ny := curr.x+dx, curr.y+dy
						if nx < 0 || ny < 0 || nx >= cols || ny >= rows || grid[ny][nx] == 0 {
							perimeter++
						} else if !visited[ny][nx] {
							visited[ny][nx] = true
							q.push(cell{x: nx, y: ny})
						}
					}
				}

				return perimeter
			}
		}
	}

	return -1
}
