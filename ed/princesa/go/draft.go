package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	pessoas := make([]int, n)

	for i := 0; i < n; i++ {
		pessoas[i] = i + 1
	}

	indEspa := e - 1


	for {

		fmt.Print("[ ")
		for i := 0; i < len(pessoas); i++ {
			fmt.Print(pessoas[i])
			if i == indEspa {
				fmt.Print(">")
			}
			fmt.Print(" ")
		}

		fmt.Println("]")

		if len(pessoas) == 1 {
			break
		}

		indVit := (indEspa + 1) % len(pessoas)
		//tirar a vitima do vetor
		pessoas = append(pessoas[:indVit], pessoas[indVit+1:]...)
		if indVit < indEspa {
			indEspa--
		}
		indEspa = (indEspa + 1) % len(pessoas)
	}
}