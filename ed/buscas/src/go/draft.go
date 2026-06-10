package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func sameStrings (stringInput []string, query []string) []int {
    count := make(map[string]int)
    for _, str := range stringInput {
        count[str]++
    }

    result := make([]int, len(query))
    for i, q := range query {
        result[i] = count[q]
    }
    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    tamanhoBuscas, _ := strconv.Atoi(scanner.Text())
    
    scanner.Scan()
    stringInput := strings.Fields(scanner.Text())
    
    scanner.Scan()
    tamanhoConsultas, _ := strconv.Atoi(scanner.Text())
    
    scanner.Scan()
    query := strings.Fields(scanner.Text())
    
    if len(stringInput) != tamanhoBuscas {
        stringInput = stringInput[:tamanhoBuscas]
    }
    
    if len(query) != tamanhoConsultas {
        query = query[:tamanhoConsultas]
    }
    
    result := sameStrings(stringInput, query)
    
    var resultStrings []string
    for _, r := range result {
        resultStrings = append(resultStrings, strconv.Itoa(r))
    }
    fmt.Println(strings.Join(resultStrings, " "))
}