package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}
	counts := make(map[int]int)
	for _, val := range vet {
		absVal := val
		if val < 0 {
			absVal = -val
		}
		counts[absVal]++
	}
	var keys []int
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var result []Pair
	for _, k := range keys {
		result = append(result, Pair{k, counts[k]})
	}
	return result
}

//SAPORRA NAO TA CONTANDO O TEMPO AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA


func teams(vet []int) []Pair {
	if len (vet) == 0 {
		return nil
	}
	var result []Pair
	currentStress := vet[0]
	count := 1
	for i := 1; i < len(vet); i++ {
		if vet[i] == currentStress{
			count++
		}else{
			result = append(result, Pair{currentStress, count})
			currentStress = vet[i]
			count = 1
		}
	}
	result = append(result, Pair{currentStress, count})
	return result
}

func mnext(vet []int) []int {
	result := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 { // Se for homem
			hasWomanNeighbor := false
			if i > 0 && vet[i-1] < 0 {
				hasWomanNeighbor = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				hasWomanNeighbor = true
			}
			if hasWomanNeighbor {
				result[i] = 1
			}
		}
	}
	return result
}

//colca 1 nos homens que nao tem mulher
func alone(vet []int) []int {
	result := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 { //é homi
			hasWomanNeighbor := false
			if i > 0 && vet[i-1] < 0 {
				hasWomanNeighbor = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				hasWomanNeighbor = true
			}
			if !hasWomanNeighbor {
				result[i] = 1
			}
		}
		
	}
	return result
}

//quantos casais podem ser formados

func couple(vet []int) int {
	menCount := make(map[int]int)
	womenCount := make(map[int]int)
	for _, val := range vet {
		if val > 0 {
			menCount[val]++
		} else if val < 0 {
			womenCount[-val]++
		}
	}
	couples := 0
	for stress, mCount := range menCount {
		if wCount, exists := womenCount[stress]; exists {
			if mCount < wCount {
				couples += mCount
			} else {
				couples += wCount
			}
		}
	}
	return couples
}
 
func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true	
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	toRemove := make(map[int]bool)
	for _, pos := range posList {
		toRemove[pos] = true
	}
	var result []int
	for i, val := range vet {
		if !toRemove[i] {
			result = append(result, val)
		}
	}
	return result
}

func clear(vet []int, value int) []int {
	var result []int
	for _, val := range vet {
		if val != value {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
