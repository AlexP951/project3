package main

import "testing"

func TestPuzzle1(t *testing.T) {
	grid, _ := ReadPuzzle("txt/sudoku-test1.txt")
	expected, _ := ReadPuzzle("txt/sudoku-test1-solved.txt")

	Solve(grid)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] != expected[r][c] {
				t.Errorf("Mismatch at (%d,%d): got %d, expected %d",
					r, c, grid[r][c], expected[r][c])
			}
		}
	}
}

func TestPuzzle2(t *testing.T) {
	grid, _ := ReadPuzzle("txt/sudoku-test2.txt")
	expected, _ := ReadPuzzle("txt/sudoku-test2-solved.txt")

	Solve(grid)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] != expected[r][c] {
				t.Errorf("Mismatch at (%d,%d): got %d, expected %d",
					r, c, grid[r][c], expected[r][c])
			}
		}
	}
}

func TestImpossible(t *testing.T) {
	grid, _ := ReadPuzzle("txt/sudoku-impossible.txt")

	if Solve(grid) {
		t.Errorf("Impossible puzzle incorrectly solved!")
	}
}
