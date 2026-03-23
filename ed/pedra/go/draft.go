package main

import "fmt"

// lançamento da pedra deve ser acima de 10m, caso falhe sera desclassificado

func main() {
	var n int
	fmt.Scan(&n)

	vencedor := -1
	difmin := 101 // Valor maior que qualquer diferença possível (máx 100)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		// ambas devem ter pelo menos 10m
		if a >= 10 && b >= 10 {
			dif := a - b
			if dif < 0 {
				dif = -dif
			}

			// Se for o primeiro competidor válido ou tiver uma pontuação melhor
			// O "<" garante que, em caso de empate, o primeiro (menor índice) permaneça
			if vencedor == -1 || dif < difmin {
				difmin = dif
				vencedor = i
			}
		}
	}

	if vencedor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(vencedor)
	}
}
