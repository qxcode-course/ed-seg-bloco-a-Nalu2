package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    semCasal := make(map[int]int)
    Tcasais := 0
    for i := 0; i < n; i++ {
        var animal int
        fmt.Scan(&animal)
        par:= -animal

        if semCasal[par] > 0 {
            Tcasais++
            semCasal[par]--
        } else {
            semCasal[animal]++
        }
    }
    fmt.Println(Tcasais)
}
