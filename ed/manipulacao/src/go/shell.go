package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var men []int
	for _, x := range vet {
		if x > 0 {
			men = append(men, x)
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	var calm []int
	for _, x := range vet {
		if x < 0 && x > -10 {
			calm = append(calm, x)
		}
	}
	return calm
}

func sortVet(vet []int) []int {
	res := append([]int{}, vet...)
	sort.Ints(res)
	return res	
}

func sortStress(vet []int) []int {
	res := append([]int{}, vet...)
	sort.Slice(res, func(i, j int) bool {
		absI := res[i]
		if absI < 0 {
			absI *= -1
		}
		absJ := res[j]
		if absJ < 0 {
			absJ *= -1
		}
		return absI < absJ
	})
	return res
}

func reverse(vet []int) []int {
	res := make([]int, len(vet))
	for i, j := 0, len(vet)-1; j >= 0; i, j = i+1, j-1 {
		res[i]= vet[j]
	}
	return res
}

func unique(vet []int) []int {
	var res []int
	exists := make(map[int]bool)
	for _, x := range vet {
		if !exists[x] {
			exists[x] = true
			res = append(res, x)
		}
	}
	return res
}

func repeated(vet []int) []int {
	var res []int
	vistos := make(map[int]int)
	for _, x := range vet {
		if vistos[x] > 0 {
			res = append(res, x)
		}
		vistos[x]++
	}
	sort.Ints(res)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

