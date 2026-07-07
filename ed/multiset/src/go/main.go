package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}
	newData := make([]int, 0, ms.capacity)
	newData = append(newData, ms.data...)
	ms.data = newData
}

func (ms *MultiSet) search(value int) (bool, int) {
	low, high := 0, ms.size-1
	for low <= high {
		mid := (low + high) / 2
		if ms.data[mid] >= value {
			high = mid - 1
		} else {
			low = mid + 1
		} 
	}
	if low < ms.size && ms.data[low] == value {
		return true, low
	}
	return false, low
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	_, pos := ms.search(value)
	ms.data = append(ms.data, 0)
	// desloca para a direita
	for i := ms.size; i > pos; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[pos] = value
	ms.size++
}

func (ms *MultiSet) Erase(value int) error {
	found, pos := ms.search(value)
	if !found {
		return fmt.Errorf("value not found")
	}
	// desloca para a esquerda
	for i := pos; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}

func (ms *MultiSet) Count(value int) int {
	found, pos := ms.search(value)
	if !found {
		return 0
	}
	count := 0
	for i := pos; i < ms.size && ms.data[i] == value; i++ {
		count++
	}
	return count
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	distinct := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			distinct++
		}
	}
	return distinct
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < ms.size; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", ms.data[i]))
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		fmt.Println("$" + strings.Join(args, " "))
		cmd := args[0]

		switch cmd {
		case "end":
			return
		case "init":
			if len(args) < 2 {
				continue
			}
			cap, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(cap)
		case "insert":
			if ms == nil {
				continue
			}
			for _, s := range args[1:] {
				value, _ := strconv.Atoi(s)
				ms.Insert(value)
			}
		case "show":
			if ms == nil {
				fmt.Println("[]")
			} else {
				fmt.Println(ms.String())
			}
		case "erase":
			if ms == nil || len(args) < 2 {
				continue
			}
			value, _ := strconv.Atoi(args[1])
			if err := ms.Erase(value); err != nil {
				fmt.Println(err.Error())
			}
		case "contains":
			if ms == nil || len(args) < 2 {
				continue
			}
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			if ms == nil || len(args) < 2 {
				continue
			}
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			if ms == nil {
				fmt.Println(0)
			} else {
				fmt.Println(ms.Unique())
			}
		case "clear":
			if ms != nil {
				ms.Clear()
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}