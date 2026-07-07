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
	n, _ := strconv.Atoi(parts[0])
	e, _ := strconv.Atoi(parts[1])
	f, _ := strconv.Atoi(parts[2])

	participar := make([]int, n)
	curr := f
	for i := 0; i < n; i++ {
		participar[i] = (i + 1) * curr
		curr *= -1
	}

	sword := -1
	for i, val := range participar {
		absVal := val
		if absVal < 0 {
			absVal = -absVal
		}
		if absVal == e {
			sword = i
			break
		}
	}

	for len(participar) > 0 {
		printState(participar, sword)
		if len(participar) == 1 {
			break
		}

		swordMan := participar[sword]
		var target int
		if swordMan > 0 {
			target = (sword + 1) % len(participar)
		} else {
			target = (sword - 1 + len(participar)) % len(participar)
		}

		participar = append(participar[:target], participar[target+1:]...)

		if swordMan > 0 {
			if target < sword {
				sword--
			}
			sword = target % len(participar)
		} else {
			if target < sword {
				sword--
			}
			sword = (target - 1 + len(participar)) % len(participar)
		}
	}
}

func printState(list []int, swordIdx int) {
	res := make([]string, len(list))
	for i, val := range list {
		if i == swordIdx {
			if val > 0 {
				res[i] = fmt.Sprintf("%d>", val)
			} else {
				res[i] = fmt.Sprintf("<%d", val)
			}
		} else {
			res[i] = fmt.Sprintf("%d", val)
		}
	}
	fmt.Printf("[ %s ]\n", strings.Join(res, " "))
}