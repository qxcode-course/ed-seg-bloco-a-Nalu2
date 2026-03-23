package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	//ainda tenho dificuldade com map
	descasados := make(map[int]int)
	totalCasais := 0

	for i := 0; i < n; i++ {
		var animal int
		fmt.Scan(&animal)

		// O par de um animal 'x' é sempre '-x'
		parBuscado := -animal

		// Se o par já estiver na lista de descasados
		if descasados[parBuscado] > 0 {
			totalCasais++
			descasados[parBuscado]-- // Remove um animal da lista pois ele formou par
		} else {
			// Se não tem par, esse animal entra para a lista de espera
			descasados[animal]++
		}
	}

	fmt.Println(totalCasais)
}
