package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, 0, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	}else {
		ms.capacity *= 2
	}
	newData := make([]int, ms.capacity)
	for i := 0; i < ms.size; i++ {
		newData[i] = ms.data[i]
	}
	ms.data = newData
}

func (ms *MultiSet) search(value int) int (bool, int) {
	low, high := 0, ms.size-1
	var last int = -1
	for low <= high {
		mid := (low + high) / 2
		if ms.data[mid] == value {
			last = mid
			low = mid + 1
		} else if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if last != -1 {
		return true, last
	}
	return false, low
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	// ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			// value, _ := strconv.Atoi(args[1])
			// ms = NewMultiSet(value)
		case "insert":
			// for _, part := range args[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
		case "erase":
			// value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
