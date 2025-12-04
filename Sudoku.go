/**
 * @author
 * Alex Pearle (alex.pearle@vanderbilt.edu)
 *
 * Sudoku solver implementation in Go.
 */

package main

/**
 * Reads a Sudoku puzzle from a filename and returns a 9×9 grid of ints.
 *
 * @param filename - name of the text file to read
 * @return 2D slice of ints representing the Sudoku puzzle
 */
func ReadPuzzle(filename string) ([][]int, error) {
	return readFileIntoGrid(filename)
}

/**
 * Attempts to solve a Sudoku puzzle.
 *
 * @param grid - 9×9 int grid
 * @return bool - true if solved, false if no solution exists
 */
func Solve(grid [][]int) bool {
	row, col, found := findEmptyCell(grid)
	if !found {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if Solve(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}

	return false
}

/**
 * Checks whether placing value n at (row, col) violates Sudoku rules.
 */
func isValid(grid [][]int, row, col, n int) bool {

	// Row
	for i := 0; i < 9; i++ {
		if grid[row][i] == n {
			return false
		}
	}

	// Column
	for i := 0; i < 9; i++ {
		if grid[i][col] == n {
			return false
		}
	}

	// 3×3 Block
	r0 := (row / 3) * 3
	c0 := (col / 3) * 3

	for r := r0; r < r0+3; r++ {
		for c := c0; c < c0+3; c++ {
			if grid[r][c] == n {
				return false
			}
		}
	}

	return true
}

/**
 * Finds the next empty cell (returns row, col, foundFlag).
 */
func findEmptyCell(grid [][]int) (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return -1, -1, false
}
