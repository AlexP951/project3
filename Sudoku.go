/**
* Alex Pearle (alexander.c.pearle@vanderbilt.edu)
* Isaac Schiffer (isaac.s.schiffer@vanderbilt.edu)
* Brantley Payne (brantley.k.payne@vanderbilt.edu)
 */

package main

const size = 9
/**
 * Attempts to solve a Sudoku puzzle.
 *
 * @param board - 9×9 int grid
 * @return bool - true if solved, false if no solution exists
 */
func Solve(board [][]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				for k := 1; k <= size; k++ {
					if check(i, j, k, board) {
						board[i][j] = k
						if Solve(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
/**
 * Checks whether placing value n at (r,c) is possible
 *
 * @return bool - if this value can be placed at the point
 */
func check(row, col, digit int, board [][]int) bool {

	for i := 0; i < size; i++ {
		if board[row][i] == digit {
			return false
		}
	}
	// Column check
	for i := 0; i < size; i++ {
		if board[i][col] == digit {
			return false
		}
	}
	// 3x3 box check
	box_row := (row / 3) * 3
	box_col := (col / 3) * 3
	for i := box_row; i < box_row+3; i++ {
		for j := box_col; j < box_col+3; j++ {
			if board[i][j] == digit {
				return false
			}
		}
	}
	return true
}
