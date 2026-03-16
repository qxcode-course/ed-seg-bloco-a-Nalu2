package main

import "fmt"

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)
	if D == -1 {
		for i := 0; i < 16; i++ {
			F--
			if F < 0 {
				F = 15
			}
			if F == P {
				fmt.Println("N")
				break
			}
			if F == H {
				fmt.Println("S")
				break
			}

		}
	} else {
		for i := 0; i < 16; i++ {
			F++
			if F > 15 {
				F = 0
			}
			if F == P {
				fmt.Println("N")
				break
			}
			if F == H {
				fmt.Println("S")
				break
			}

		}

	}
}
