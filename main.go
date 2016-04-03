package main

import (
	"fmt"
	"time"
)

var sudoku1 = [9][9]int{
	{0, 0, 6, 0, 0, 8, 5, 0, 0},
	{0, 0, 0, 0, 7, 0, 6, 1, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 9},

	{0, 0, 0, 0, 9, 0, 0, 0, 1},
	{0, 0, 1, 0, 0, 0, 8, 0, 0},
	{4, 0, 0, 5, 3, 0, 0, 0, 0},

	{1, 0, 7, 0, 5, 3, 0, 0, 0},
	{0, 5, 0, 0, 6, 4, 0, 0, 0},
	{3, 0, 0, 1, 0, 0, 0, 6, 0},
}

var sudoku2 = [9][9]int{
	{0, 0, 0, 2, 0, 4, 8, 1, 0},
	{0, 4, 0, 0, 0, 8, 2, 6, 3},
	{3, 0, 0, 1, 6, 0, 0, 0, 4},

	{1, 0, 0, 0, 4, 0, 5, 8, 0},
	{6, 3, 5, 8, 2, 0, 0, 0, 7},
	{2, 0, 0, 5, 9, 0, 1, 0, 0},

	{9, 1, 0, 7, 0, 0, 0, 4, 0},
	{0, 0, 0, 6, 8, 0, 7, 0, 1},
	{8, 0, 0, 4, 0, 3, 0, 5, 0},
}

var sudoku3 = [9][9]int{
	{0, 0, 6, 0, 0, 8, 5, 0, 0},
	{0, 0, 0, 0, 7, 0, 6, 1, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 9},

	{0, 0, 0, 0, 9, 0, 0, 0, 1},
	{0, 0, 1, 0, 0, 0, 8, 0, 0},
	{4, 0, 0, 5, 3, 0, 0, 0, 0},

	{1, 0, 7, 0, 5, 3, 0, 0, 0},
	{0, 5, 0, 0, 6, 4, 0, 0, 0},
	{3, 0, 0, 1, 0, 0, 0, 6, 0},
}

var sudoku4 = [9][9]int{
	{8, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 3, 6, 0, 0, 0, 0, 0},
	{0, 7, 0, 0, 9, 0, 2, 0, 0},

	{0, 5, 0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 4, 5, 7, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 3, 0},

	{0, 0, 1, 0, 0, 0, 0, 6, 8},
	{0, 0, 8, 5, 0, 0, 0, 1, 0},
	{0, 9, 0, 0, 0, 0, 4, 0, 0},
}

type Unknown struct {
	x int
	y int
}

func main() {
	solveSudoku(&sudoku1)
	solveSudoku(&sudoku2)
	solveSudoku(&sudoku3)
	solveSudoku(&sudoku4)
}

func solveSudoku(puzzle *[9][9]int) {
	start := time.Now()
	unknowns := make([]Unknown, 0, 81)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if puzzle[i][j] == 0 {
				unknowns = append(unknowns, Unknown{i, j})
			}
		}
	}
	cur := 0
	for cur < len(unknowns) {
		if solve(puzzle, cur, unknowns) {
			cur++
		} else {
			cur--
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Sudoku solver took %s\n", elapsed)
	for i := range puzzle {
		fmt.Println(puzzle[i])
	}

}

func solve(puzzle *[9][9]int, cur int, unknowns []Unknown) bool {
	unk := unknowns[cur]
	value := puzzle[unk.x][unk.y] + 1
	for i := value; i <= 9; i++ {
		valid := true
		for j := 0; j < 9; j++ {
			if puzzle[unk.x][j] == i || puzzle[j][unk.y] == i || puzzle[(int(unk.x/3)*3)+(j/3)][(int(unk.y/3)*3)+(j%3)] == i {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		puzzle[unk.x][unk.y] = i
		return true
	}
	puzzle[unk.x][unk.y] = 0
	return false
}
