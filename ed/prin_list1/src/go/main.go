package main

import (
	"fmt"
	"strings"
)

func ToStr(l *DList[int], sword *DNode[int]) string {
	var values []string
	for n := l.root.next; n != l.root; n = n.next {
		if n == sword {
			values = append(values, fmt.Sprintf("%d>", n.Value))
		} else {
			values = append(values, fmt.Sprint(n.Value))
		}
	}
	return "[ " + strings.Join(values, " ") + " ]"
}

func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil {
		return nil
	}
	proximo := it.next
	if proximo == l.root {
		proximo = proximo.next
	}
	return proximo
}

func main() {
	var qtd, chosen int
	if _, err := fmt.Scan(&qtd, &chosen); err != nil {
		return
	}

	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	sword := l.Front()
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}

	for i := 0; i < qtd-1; i++ {
		fmt.Println(ToStr(l, sword))
		morreu := Next(l, sword)
		l.Erase(morreu)
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}