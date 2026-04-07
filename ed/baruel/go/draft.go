package main
import "fmt"

func main() {
    //perfeito exemplo de utilização de map, foda que eu nao sei usar direito
    var qtdT, qtdP int
    fmt.Scan(&qtdT, &qtdP)
    contagem := make(map[int]int)
    temRep := false
    priRep := true

    //ler as figurinhas
    for i := 0; i < qtdP; i++ {
        var fig int
        fmt.Scan(&fig)
        //se valor for maior que 0, tem repetida
        if contagem[fig] > 0 {
            if !priRep{
                fmt.Print(" ")
            }
            fmt.Print(fig)
            priRep = false
            temRep = true
        }
        contagem[fig]++
    }
    
    if !temRep {
        fmt.Print("N")
    }
    fmt.Println("")

    //quais faltam

    temFalta := false
    priFalta := true

    for i := 1; i <= qtdT; i++ {
        if contagem[i] == 0 {
            if !priFalta{
                fmt.Print(" ")
            }
            fmt.Print(i)
            priFalta = false
            temFalta = true
        }
    }

    if !temFalta {
        fmt.Print("N")
    }
    fmt.Println("")

    
}
