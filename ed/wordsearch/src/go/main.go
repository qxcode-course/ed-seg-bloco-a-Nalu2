package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}
	m := len(grid)
	n := len(grid[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if dfs(grid, r, c, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(grid [][]byte, r, c int, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}
	temp := grid[r][c]
	grid[r][c] = 0 

	found := dfs(grid, r+1, c, word, index+1) ||
		dfs(grid, r-1, c, word, index+1) ||
		dfs(grid, r, c+1, word, index+1) ||
		dfs(grid, r, c-1, word, index+1)
	
	grid[r][c] = temp
	return found
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
