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
	"strings"
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter puzzle filename (Press Enter for sudoku-test1.txt): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		input = "txt/sudoku-test1.txt"
	} else {
		input = "txt/" + input
	}

	grid, err := ReadPuzzle(input)
	if err != nil {
		fmt.Println("Error reading puzzle:", err)
		return
	}

	fmt.Println("\nInitial puzzle:")
	PrintPuzzle(grid)

	solved := Solve(grid)

	if solved {
		fmt.Println("\nSolved puzzle:")
		PrintPuzzle(grid)
	} else {
		fmt.Println("This puzzle has no solution.")
	}
}
