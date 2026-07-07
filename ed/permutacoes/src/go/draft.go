package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func nextPermutation(s []rune) bool {
	n := len(s)
	i := n - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for s[j] <= s[i] {
		j--
	}
	s[i], s[j] = s[j], s[i]
	// reverse suffix
	left, right := i+1, n-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	s := strings.TrimSpace(scanner.Text())
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	fmt.Println(string(chars))
	for nextPermutation(chars) {
		fmt.Println(string(chars))
	}
}