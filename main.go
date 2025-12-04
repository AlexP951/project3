/**
*
* Alex Pearle (alexander.c.pearle@vanderbilt.edu)
* Isaac Schiffer (isaac.s.schiffer@vanderbilt.edu)
* Brantley Payne (brantley.k.payne@vanderbilt.edu)
*
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

/**
 * Prints a Sudoku puzzle with borders.
 *
 * @param grid - the 9×9 sudoku grid to print
 */
func PrintSudoku(grid [][]int) {
	line := "+-------+-------+-------+"
	for i := 0; i < 9; i++ {
		if i % 3 == 0 {
			fmt.Println(line)
		}
		for j := 0; j < 9; j++ {
			if j % 3 == 0 {
				fmt.Print("| ")
			}
			if grid[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Println("|")
	}
	fmt.Println(line)
}

/**
 * Main program:
 * - prompts user for puzzle filename
 * - loads puzzle
 * - prints before + after
 */
func main() {
	var filename string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter puzzle text file (assumes file is in \"txt\" folder).")
	fmt.Println("Pressing <Enter> will run the file \"sudoku-test1.txt\".")
	if scanner.Scan() {
		filename = scanner.Text()
	}
	if filename == "" {
		filename = "sudoku-test1.txt"
	}
	path := filepath.Join("txt", filename)
	puzzle := readFileIntoGrid(path)
	fmt.Println("\nPuzzle:\n")
	PrintSudoku(puzzle)
	fmt.Println()
	solved := Solve(puzzle)
	if solved {
		fmt.Println("Solution:\n")
		PrintSudoku(puzzle)
		fmt.Println()
	} else {
		fmt.Println("This puzzle has no solution.\n")
	}

}
