package main
import "fmt"

//precisa de uma estrutura para os pontos em cartesiano
//aqui foi gpt mesmo viu
type Ponto struct{
    x, y int
}

func main() {
    var q int
    var d string
    fmt.Scan(&q, &d)
    cobra := make([]Ponto, q)

    for i := 0; i < q; i++ {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    newHead := cobra[0]
    switch d {
        case "D":
            newHead.y++
        case "U":
            newHead.y--
        case "R":
            newHead.x++
        case "L":
            newHead.x--
    }

    fmt.Println(newHead.x, newHead.y)

    for i := 0; i < q-1; i++ {
        fmt.Println(cobra[i].x, cobra[i].y)
    }


}
