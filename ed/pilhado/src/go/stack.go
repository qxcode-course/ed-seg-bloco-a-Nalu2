package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	header := strings.TrimSpace(strings.TrimPrefix(scanner.Text(), "\ufeff"))
	var n, m int
	if _, err := fmt.Sscanf(header, "%d %d", &n, &m); err != nil || n <= 0 || m <= 0 {
		return
	}

	maze := make([][]rune, n)
	var inicio, fim Pos

	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		runes := make([]rune, m)

		for j, ch := range line {
			if j < m {
				runes[j] = ch
			}
		}
		for j := len(line); j < m; j++ {
			runes[j] = ' '
		}
		maze[i] = runes

		for j, ch := range runes {
			if ch == 'I' {
				inicio = Pos{i, j}
			} else if ch == 'F' {
				fim = Pos{i, j}
			}
		}
	}

	visitado := make([][]bool, n)
	for i := range visitado {
		visitado[i] = make([]bool, m)
	}

	caminho := NewStack[Pos]()
	caminho.Push(inicio)
	visitado[inicio.l][inicio.c] = true

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for !caminho.IsEmpty() {
		atual := caminho.Top()

		if atual == fim {
			break
		}

		validos := []Pos{}
		for _, d := range dirs {
			linha := atual.l + d[0]
			coluna := atual.c + d[1]
			if linha < 0 || linha >= n || coluna < 0 || coluna >= m {
				continue
			}
			if maze[linha][coluna] == '#' {
				continue
			}
			if visitado[linha][coluna] {
				continue
			}
			validos = append(validos, Pos{linha, coluna})
		}

		if len(validos) > 0 {
			proximo := validos[0]
			caminho.Push(proximo)
			visitado[proximo.l][proximo.c] = true
		} else {
			caminho.Pop()
		}
	}
	if !caminho.IsEmpty() {
		for _, p := range caminho.data {
			if maze[p.l][p.c] != '#' {
				maze[p.l][p.c] = '.'
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(string(maze[i]))
	}
}