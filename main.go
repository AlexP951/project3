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
func PrintPuzzle(grid [][]int) {
	fmt.Println("+-------+-------+-------+")
	for r := 0; r < 9; r++ {
		fmt.Print("| ")
		for c := 0; c < 9; c++ {
			fmt.Printf("%d ", grid[r][c])
			if (c+1)%3 == 0 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (r+1)%3 == 0 {
			fmt.Println("+-------+-------+-------+")
		}
	}
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
