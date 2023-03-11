package main

import (
	"fmt"
)

// 19% ？

func main() {
	var rows, cols, cost int
	fmt.Scanf("%d%d%d\n", &rows, &cols, &cost)

	board := make([][]cell, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]cell, cols)
	}

	for y := 0; y < rows; y++ {
		var colors string
		fmt.Scanln(&colors)
		for x, color := range colors {
			board[y][x].color = byte(color)
		}
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			fmt.Scanf("%d", &(board[y][x].gold))
		}
	}

	fmt.Println(solve(board, cost))
}

type cell struct {
	gold  int
	color byte
}

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func solve(board [][]cell, cost int) int {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	maxFinalGolds := 0
	visited[0][0] = true
	recursive(board, cost, rows, cols, 0, 0, visited, 0, &maxFinalGolds)

	return maxFinalGolds
}

func recursive(board [][]cell, cost int, rows int, cols int, x, y int, visited [][]bool, golds int, maxFinalGolds *int) {
	golds += board[y][x].gold
	if golds > *maxFinalGolds {
		*maxFinalGolds = golds
	}
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		xx, yy := x+dx, y+dy
		if yy < 0 || yy >= rows || xx < 0 || xx >= cols {
			continue
		}

		if visited[yy][xx] {
			continue
		}

		target := &board[yy][xx]
		current := &board[y][x]
		consumed := 0
		if current.color != target.color {
			if golds < cost {
				continue
			}
			consumed = cost
		}

		visited[yy][xx] = true
		recursive(board, cost, rows, cols, xx, yy, visited, golds-consumed, maxFinalGolds)
		// visited[yy][xx] = false
	}
}
