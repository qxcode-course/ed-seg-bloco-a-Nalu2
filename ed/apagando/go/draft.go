package main

import "fmt"

func main() {
	var n, ret int
	fmt.Scan(&n)
	var qtd = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&qtd[i])
	}

	fmt.Scan(&ret)
	var num = make([]int, ret)
	for i := 0; i < ret; i++ {
		fmt.Scan(&num[i])
	}

	//aprendendo a usar mapa

	mapa := make(map[int]bool)

	for i := 0; i < ret; i++ {
		mapa[num[i]] = true
	}

	//isso seria um novo slice com os valores filtrados

	result := []int{}

	for i := 0; i < n; i++ {
		if !mapa[qtd[i]] { // se nao ta no mapa mantem
			result = append(result, qtd[i])
		}
	}
	///o mapa tem [] inicio e fim, tenho que retirar para essa atividade, cara aqui terminou com uma cambiara ferrada

	for i := 0; i < len(result); i++ {
		fmt.Print(result[i], " ")

	}
	fmt.Println()
}
