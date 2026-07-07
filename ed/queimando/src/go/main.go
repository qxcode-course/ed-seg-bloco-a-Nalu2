package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct{ l, c int }

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("empty stack")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func burnTrees(grid [][]rune, l, c int) {
	nl := len(grid)
	if nl == 0 {
		return
	}
	nc := len(grid[0])
	if l < 0 || l >= nl || c < 0 || c >= nc {
		return
	}

	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.Empty() {
		p := stack.Pop()
		if p.l < 0 || p.l >= nl || p.c < 0 || p.c >= nc {
			continue
		}
		if grid[p.l][p.c] != '#' {
			continue
		}
		grid[p.l][p.c] = 'o'

		stack.Push(Pos{p.l - 1, p.c})
		stack.Push(Pos{p.l + 1, p.c})
		stack.Push(Pos{p.l, p.c - 1})
		stack.Push(Pos{p.l, p.c + 1})
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	var nl, nc, lfire, cfire int
	fmt.Sscan(scanner.Text(), &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			grid[i] = []rune(strings.Repeat(".", nc))
			continue
		}
		line := scanner.Text()
		if len(line) < nc {
			line += strings.Repeat(".", nc-len(line))
		}
		if len(line) > nc {
			line = line[:nc]
		}
		grid[i] = []rune(line)
	}
	if lfire >= 0 && lfire < nl && cfire >= 0 && cfire < nc {
		burnTrees(grid, lfire, cfire)
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}