package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Reserve(newCapacity int) {
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

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		newCap := v.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		v.Reserve(newCap)
	}
	v.data[v.size] = value
	v.size++
}

func (v *Vector) PopBack() (int, error) {
	if v.size == 0 {
		return 0, fmt.Errorf("vector is empty")
	}
	v.size--
	return v.data[v.size], nil
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) String() string {
	if v.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < v.size; i++ {
		sb.WriteString(strconv.Itoa(v.data[i]))
		if i < v.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (v *Vector) Get(index int) int {
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

func (v *Vector) Clear() {
	v.size = 0
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}
	if v.size == v.capacity {
		newCap := v.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		v.Reserve(newCap)
	}
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}

func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector) Slice(start int, end int) *Vector {
	if v.size == 0 {
		return NewVector(0)
	}

	resolveIndex := func(idx, size int) int {
		if idx < 0 {
			idx = size + idx
			if idx < 0 {
				return 0
			}
			return idx
		}
		if idx > size {
			return size
		}
		return idx
	}

	actualStart := resolveIndex(start, v.size)
	actualEnd := resolveIndex(end, v.size)

	if actualStart > actualEnd {
		actualStart = actualEnd
	}

	subSize := actualEnd - actualStart
	subCap := v.capacity - actualStart

	return &Vector{
		data:     v.data[actualStart:actualEnd],
		size:     subSize,
		capacity: subCap,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var v *Vector

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			if len(parts) < 2 {
				fmt.Println("fail: falta argumento")
				continue
			}
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			if v == nil {
				continue
			}
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			if v == nil {
				fmt.Println("[]")
				continue
			}
			fmt.Println(v.String())
		case "status":
			if v == nil {
				fmt.Println("size:0 capacity:0")
				continue
			}
			fmt.Println(v.Status())
		case "pop":
			if v == nil {
				fmt.Println("vector is empty")
				continue
			}
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			if v == nil {
				fmt.Println("index out of range")
				continue
			}
			if len(parts) < 3 {
				fmt.Println("index out of range")
				continue
			}
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			if v == nil {
				fmt.Println("index out of range")
				continue
			}
			if len(parts) < 2 {
				fmt.Println("index out of range")
				continue
			}
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			if v == nil {
				fmt.Println(-1)
				continue
			}
			if len(parts) < 2 {
				fmt.Println(-1)
				continue
			}
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.IndexOf(value))
		case "contains":
			if v == nil {
				fmt.Println("false")
				continue
			}
			if len(parts) < 2 {
				fmt.Println("false")
				continue
			}
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			if v != nil {
				v.Clear()
			}
		case "capacity":
			if v == nil {
				fmt.Println(0)
				continue
			}
			fmt.Println(v.Capacity())
		case "get":
			if v == nil {
				fmt.Println("index out of range")
				continue
			}
			if len(parts) < 2 {
				fmt.Println("index out of range")
				continue
			}
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			if v == nil {
				fmt.Println("index out of range")
				continue
			}
			if len(parts) < 3 {
				fmt.Println("index out of range")
				continue
			}
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "reserve":
			if v == nil {
				continue
			}
			if len(parts) < 2 {
				continue
			}
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			if v == nil {
				fmt.Println("[]")
				continue
			}
			if len(parts) < 3 {
				fmt.Println("[]")
				continue
			}
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice.String())
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}