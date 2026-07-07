package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sameStrings(stringInput []string, query []string) []int {
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

	if !scanner.Scan() {
		return
	}
	tamanhoBuscas, _ := strconv.Atoi(scanner.Text())

	if !scanner.Scan() {
		return
	}
	stringInput := strings.Fields(scanner.Text())

	if !scanner.Scan() {
		return
	}
	tamanhoConsultas, _ := strconv.Atoi(scanner.Text())

	if !scanner.Scan() {
		return
	}
	query := strings.Fields(scanner.Text())

	if len(stringInput) > tamanhoBuscas {
		stringInput = stringInput[:tamanhoBuscas]
	}
	if len(query) > tamanhoConsultas {
		query = query[:tamanhoConsultas]
	}

	result := sameStrings(stringInput, query)

	var resultStrings []string
	for _, r := range result {
		resultStrings = append(resultStrings, strconv.Itoa(r))
	}
	fmt.Println(strings.Join(resultStrings, " "))
}