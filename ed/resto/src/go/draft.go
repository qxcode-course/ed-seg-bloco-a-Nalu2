package main
import "fmt"
func resto(n int) {
    if n == 0{
        return
    }
    resto(n/2)
    fmt.Println(n/2, n%2)
}
func main() {
    var n int
    fmt.Scan(&n)
    resto(n)
}