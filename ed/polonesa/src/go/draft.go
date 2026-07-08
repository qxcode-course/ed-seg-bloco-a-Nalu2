package main
import (
    "strings"
    "fmt"
    "os"
    "bufio"
)

func precedente(op string) int{
    switch op{
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    case "^":
        return 3
    }
    return 0
    
}

func isOperator (token string) bool {
    return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func ShutingYard(input string) string {
    tokens := strings.Split(input, " ")
    var output []string
    var stack []string

    for _, token := range tokens {
        if token == ""{
            continue 
        }

        if isOperator(token) {
            for len(stack) > 0 && precedente (stack[len(stack)-1]) >= precedente(token){
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, token)
        } else {
            output= append(output, token)
        }
    }

    for len(stack) > 0 {
        output = append(output, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }

    return strings.Join(output, " ")

}


func main(){
    scanner:= bufio.NewScanner(os.Stdin)
    if scanner.Scan(){
        input := scanner.Text()
        res :=ShutingYard(input)
        fmt.Println(res)
    }
}