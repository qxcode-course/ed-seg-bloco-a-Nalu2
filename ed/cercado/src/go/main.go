package main

import (
	"bufio"
	"fmt"
	"os"
)

func dsf(board[][]byte, r, c int){
	if r<0 || r >= len(board) || c <0 || c>= len(board[0]) || board[r][c] != 'O'{
		return
	}
	board[r][c] ='E'
	dsf(board, r+1, c)
	dsf(board, r-1, c)
	dsf(board, r, c+1)
	dsf(board, r, c-1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	nrows := len(board)
	ncols := len(board[0])

	for i :=0; i <nrows; i++{
		dsf(board, i, 0)
		dsf(board, i, ncols-1)
	}
	for j :=0; j <ncols; j++{
		dsf(board, 0, j)
		dsf(board, nrows-1, j)
	}
	for i:=0; i <nrows; i++ {
		for j := 0; j<ncols; j++{
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'E'{
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
