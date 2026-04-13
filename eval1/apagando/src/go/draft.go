package main
import "fmt"
func main() {
	var n, ret int
	fmt.Scan(&n)
	var qtd = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&qtd[i])
	}
	//aqui da pra fazer uma melhoria legaaal de memoria
	fmt.Scan(&ret)
	mapa := make(map[int]bool)
	for i := 0; i < ret; i++ {
		var aux int
		fmt.Scan(&aux)
		mapa[aux] = true
	}

	res :=[]int{}
	for _, v := range qtd {
		if !mapa[v] {
			res = append(res, v)
		}
	}

	for i, v := range res {
		fmt.Print(v)
		if i < len(res)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println(" ")



}
