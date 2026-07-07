package main
import (
	"bufio"
	"fmt"
	"os"
)

func balanced (s string) bool {
    stack := []rune{}
    for _, char := range s {
        switch char {
        case '(', '[':
            stack = append(stack, char)
        case ')':
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
        case ']':
            if len(stack) == 0 || stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    expr := scanner.Text()
    if balanced(expr) {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}

