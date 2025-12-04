package main

import (
	"bufio"
	"os"
	"strconv"
)

/**
 * Reads a text file containing a 9×9 Sudoku puzzle
 *
 * @param filename text file containing 9 rows of 9 numbers
 * @return 2D int
 */
func readFileIntoGrid(filename string) [][]int {
	// opens the file and returns error if doesn't work
	file, _ := os.Open(filename)
	// makes sure file is closed properly
	defer file.Close()

	grid := make([][]int, 0, 9)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		nums := strings.Fields(line)

		// convert strings into integers
		row := make([]int, 9)
		for i, n := range nums {
			val, _ := strconv.Atoi(n)
			row[i] = val
		}
		grid = append(grid, row)
	}
	return grid
}
