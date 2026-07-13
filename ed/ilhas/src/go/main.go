package main

import (
	"bufio"
	"fmt"
	"os"
)

func dsf(grid [][]byte, r, c int){
	if r<0 || r >= len(grid) || c <0 || c>= len(grid[0]) || grid[r][c] == '0'{
		return
	}
	grid[r][c] ='0'
	dsf(grid, r+1, c)
	dsf(grid, r-1, c)
	dsf(grid, r, c+1)
	dsf(grid, r, c-1)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0{
		return 0
	}
	count := 0
	for r := 0; r< len(grid); r ++{
		for c :=0; c < len(grid[0]); c++ {
			if grid[r][c] == '1'{
				count ++
				dsf(grid,r,c)
			}
		}
	}
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
