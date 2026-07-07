package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	t, _ := strconv.Atoi(parts[0])
	r, _ := strconv.Atoi(parts[1])

	if !scanner.Scan() {
		return
	}
	vec := strings.Fields(scanner.Text())

	if t == 0 {
		fmt.Println("[ ]")
		return
	}

	start := (t - (r % t)) % t
	res := make([]string, t)
	for i := 0; i < t; i++ {
		res[i] = vec[(start+i)%t]
	}
	fmt.Printf("[ %s ]\n", strings.Join(res, " "))
}