package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) size() int {
	return v.size
}

func (v *Vector) capacity() int {
	return v.capacity
}

func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		newCap := v.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		v.reserve(newCap)
	}
	v.data[v.size] = value
	v.size++
}

func (v *Vector) reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}
	newData := make([]int, newCapacity)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}
	v.data = newData
	v.capacity = newCapacity
}

func (v *Vector) PopBack() error {
	if v.size == 0 {
		return fmt.Errorf("vector is empty")
	}
	v.size--
	return nil
}

func (v *Vector) status() string {
	return fmt.Sprintf("size: %d, capacity: %d", v.size, v.capacity)
}

func (v *Vector) string() string {
	if v.size == 0 {
		return "[]"
	}
	var result strings.Builder
	result.WriteString("[")
	for i := 0; i < v.size; i++ {
		fmt.Fprintf(&result, "%d", v.data[i])
		if i < v.size-1 {
			result.WriteString(", ")
		}
	}
	result.WriteString("]")
	return result.String()
}

func (v *Vector) get(index int) int {
	return v.data[index]
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}

func (v *Vector) clear() {
	v.size = 0
}

func (v *Vector) insert(index int, value int) []int {
	if index < 0 || index > v.size {
		return errors.New("index out of range")
	}
	if v.size == v.capacity {
		newCap := v.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		v.reserve(newCap)
	}
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++
	return nil
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	// v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			// value, _ := strconv.Atoi(parts[1])
			// v = NewVector(value)
		case "push":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// 	v.PushBack(value)
			// }
		case "show":
			// fmt.Println(v)
		case "status":
			// fmt.Println(v.Status())
		case "pop":
			// err := v.PopBack()
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "insert":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Insert(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "erase":
			// index, _ := strconv.Atoi(parts[1])
			// err := v.Erase(index)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "indexOf":
			// value, _ := strconv.Atoi(parts[1])
			// index := v.IndexOf(value)
			// fmt.Println(index)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
			// if v.Contains(value) {
			// 	fmt.Println("true")
			// } else {
			// 	fmt.Println("false")
			// }
		case "clear":
			// v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			// index, _ := strconv.Atoi(parts[1])
			// value, err := v.At(index)
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(value)
			// }
		case "set":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Set(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// 
		case "reserve":
			// newCapacity, _ := strconv.Atoi(parts[1])
			// v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
