package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	nl := len(grid)
	if nl == 0 {
		return
	}
	nc := len(grid[0])
	if l < 0 || l >= nl || c < 0 || c >= nc {
		return
	}
	if grid[l][c] != 'T' {
		return
	}
	grid[l][c] = 'F'
	burnTrees(grid, l-1, c)
	burnTrees(grid, l+1, c)
	burnTrees(grid, l, c-1)
	burnTrees(grid, l, c+1)
	//errei na ordem mas agora vai
	
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
