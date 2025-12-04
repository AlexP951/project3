/**
* Alex Pearle (alexander.c.pearle@vanderbilt.edu)
* Isaac Schiffer (isaac.s.schiffer@vanderbilt.edu)
* Brantley Payne (brantley.k.payne@vanderbilt.edu)
*/

package main

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	grid := readFileIntoGrid("txt/sudoku-test1.txt")
	expected := readFileIntoGrid("txt/sudoku-test1-solved.txt")

	Solve(grid)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] != expected[r][c] {
				t.Errorf("Incorrect Solver at (%d, %d): got %d, expected %d",
					r, c, grid[r][c], expected[r][c])
			}
		}
	}
	fmt.Println("Puzzle 1 solved correctly")
}

func TestPuzzle2(t *testing.T) {
	grid := readFileIntoGrid("txt/sudoku-test2.txt")
	expected := readFileIntoGrid("txt/sudoku-test2-solved.txt")

	Solve(grid)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] != expected[r][c] {
				t.Errorf("Incorrect Solver at (%d, %d): got %d, expected %d",
					r, c, grid[r][c], expected[r][c])
			}
		}
	}
	fmt.Println("Puzzle 2 solved correctly")
}

func TestImpossible(t *testing.T) {
	grid := readFileIntoGrid("txt/sudoku-impossible.txt")

	if Solve(grid) {
		t.Errorf("Impossible puzzle incorrectly solved")
	}
	fmt.Println("Impossible puzzle did not have a solution which is correct")
}
