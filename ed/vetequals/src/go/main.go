package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EhVazio(vet []int) bool {
	return fmt.Sprint(vet) == "[]"
}

// não altere a assinatura
func equals(a []int, b []int) bool {
	if EhVazio(a) && EhVazio(b){
		return true
	}
	if EhVazio(a) || EhVazio(b) {
		return false
	}
	if a[0] != b[0]{
		return false
	}
	return equals(a[1:], b[1:])


}

func main() {
	scanner := bufio.NewScanner(os.Stdin)


	scanner.Scan()
	vet1 := str2slice(scanner.Text())
	scanner.Scan()
	vet2 := str2slice(scanner.Text())
	if equals(vet1, vet2) {
		fmt.Println("iguais")
	} else {
		fmt.Println("diferentes")
	}
}

func str2slice(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, 0)
	for i := 1; i < len(parts)-1; i++ {
		value, _ := strconv.Atoi(parts[i])
		nums = append(nums, value)
	}
	return nums
}
