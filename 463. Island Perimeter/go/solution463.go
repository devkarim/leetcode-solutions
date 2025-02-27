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

func islandPerimeter(grid [][]int) int {
	rows := len(grid)    // number of rows
	cols := len(grid[0]) // number of cols

	perimeter := 0

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 1 {
				for _, d := range directions {
					dx, dy := d[0], d[1]
					nx, ny := x+dx, y+dy
					if nx < 0 || ny < 0 || nx >= cols || ny >= rows || grid[ny][nx] == 0 {
						perimeter++
					}
				}
			}
		}
	}

	return perimeter
}
