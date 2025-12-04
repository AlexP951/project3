package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Reads a text file containing a 9×9 Sudoku puzzle.
 *
 * @param filename - text file containing 9 rows of 9 numbers
 * @return 2D int slice
 */
func readFileIntoGrid(filename string) ([][]int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grid := make([][]int, 0, 9)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		nums := strings.Fields(line)
		if len(nums) != 9 {
			return nil, fmt.Errorf("invalid puzzle row: %s", line)
		}

		row := make([]int, 9)
		for i, n := range nums {
			val, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("non-integer: %s", n)
			}
			row[i] = val
		}
		grid = append(grid, row)
	}

	if len(grid) != 9 {
		return nil, fmt.Errorf("puzzle must contain 9 rows, found %d", len(grid))
	}

	return grid, nil
}
