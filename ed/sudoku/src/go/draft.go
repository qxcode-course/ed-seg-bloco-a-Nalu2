package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	grid := make([][]rune, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		grid[i] = []rune(scanner.Text())
	}

	if solve(grid, n) {
		for _, row := range grid {
			fmt.Println(string(row))
		}
	}
}

func solve(grid [][]rune, n int) bool {
	bloco := 2
	if n == 9 {
		bloco = 3
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' {
				continue 
			}
			max := 9
			if n == 4 {
				max = 4
			}
			for num := 1; num <= max; num++ {
				c := rune('0' + num)
				if valido(grid, i, j, c, bloco) {
					grid[i][j] = c
					if solve(grid, n) {
						return true
					}
					grid[i][j] = '.'
				}
			}
			return false 
		}
	}
	return true 
}

func valido(grid [][]rune, lin, col int, c rune, bloco int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		if grid[lin][i] == c || grid[i][col] == c {
			return false
		}
	}
	l0 := (lin / bloco) * bloco
	c0 := (col / bloco) * bloco
	for i := l0; i < l0+bloco; i++ {
		for j := c0; j < c0+bloco; j++ {
			if grid[i][j] == c {
				return false
			}
		}
	}
	return true
}