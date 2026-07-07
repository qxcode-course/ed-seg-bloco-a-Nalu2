package main
import (
	"bufio"
	"fmt"
	"os"
)

type Posto struct {
    gasolina int
    distancia int
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }

    var n int
    fmt.Sscanf(scanner.Text(), "%d", &n)
    postos := make([]Posto, n)
    tGasolina, tDistancia := 0, 0

    for i := 0; i < n; i++ {
        if !scanner.Scan() {
            break
        }
        var gas, dist int
        fmt.Sscanf(scanner.Text(), "%d %d", &gas, &dist)
        postos[i] = Posto{gasolina: gas, distancia: dist}
        tGasolina += gas
        tDistancia += dist
    }

    if tGasolina < tDistancia {
        fmt.Println(-1)
        return
    }

    start := 0
    currGas := 0

    for i := 0; i < n; i++ {
        currGas += postos[i].gasolina - postos[i].distancia
        if currGas < 0 {
            start = i + 1
            currGas = 0
        }
    }

    fmt.Println(start)
}