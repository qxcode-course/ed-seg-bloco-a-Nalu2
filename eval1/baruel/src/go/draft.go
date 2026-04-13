package main
import "fmt"
func main() {
    var qtdTotal, qtdPossui int
    fmt.Scan(&qtdTotal, &qtdPossui)

    cont := make(map[int]int)
    temRep := false
    priRep := true

    //identificar repetida
    for i := 0; i < qtdPossui; i++ {
        var fig int
        fmt.Scan(&fig)
        if cont[fig] > 0 {
            if !priRep {
                fmt.Print(" ")
            }
            fmt.Print(fig)
            temRep = true
            priRep = false
        }
        cont[fig]++
    }

    if !temRep {
        fmt.Print("N")
    }
    fmt.Println("")
    //identificar o que flata

    temFalta := false
    priFalta := true
    for i := 1; i <= qtdTotal; i++ {
        if cont[i] == 0 {
            if !priFalta {
                fmt.Print(" ")
            }
            fmt.Print(i)
            temFalta = true
            priFalta = false
        }
    }

    if !temFalta {
        fmt.Print("N")
    }

    fmt.Println("")

}
