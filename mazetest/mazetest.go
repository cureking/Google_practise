package main

import (
	"fmt"
	"os"
)

func printMaze(maze [][]int) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%4d", maze[i][j])
		}
		fmt.Println()
	}
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d ", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanf(file, " ")
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func (p1 point) add(p2 point) point {
	p3 := point{p1.i + p2.i, p1.j + p2.j}
	return p3
}

func (p1 point) at(maze [][]int) (int, bool) {
	if p1.i < 0 || p1.i >= len(maze) {
		return -1, false
	}
	if p1.j < 0 || p1.j >= len(maze[p1.i]) {
		return -1, false
	}
	return maze[p1.i][p1.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)

			if val, ok := next.at(maze); !ok || val != 0 {
				continue
			}
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("./mazetest/maze.in")

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)
	printMaze(steps)
}
