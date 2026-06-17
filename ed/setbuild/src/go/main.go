package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data 	 []int
	size 	 int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int){
	if newCapacity <= s.capacity {
		return
	}
	newData := make([]int, newCapacity)
	for i := 0; i < s.size; i++{
		newData[i] = s.data[i]
	}
	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int{
	low, high := 0, s.size-1
	for low <= high {
		mid := (low + high) / 2
		if s.data[mid] == value {
			return mid
		} else if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (s *Set) findIndex(value int) int{
	low, high := 0, s.size
	for low < high {
		mid := (low + high) / 2
		if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func (s *Set) insertAt(value int, index int) error{
	if index < 0 || index > s.size {
		return fmt.Errorf("index out of range")
	}
	if s.size == s.capacity {
		if s.capacity == 0 {
			s.reserve(1)
		} else {
			s.reserve(s.capacity * 2)
		}
	}
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++
	return nil
}

func (s *Set) Insert(value int) {
    if s.binarySearch(value) != -1 {
        return
    }
    pos := s.findIndex(value)   
    s.insertAt(value, pos)      
}


func (s *Set) removeAt(index int) error{
	if index < 0 || index >= s.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}

func ( s*Set) Contains(value int) bool{
	return s.binarySearch(value) != -1
}

func (s *Set) Erase(value int) bool {
    idx := s.binarySearch(value)
    if idx == -1 {
        return false
    }
    s.removeAt(idx)  
    return true
}

func (s *Set) String() string{
	if s.size == 0 {
		return "[]"
	}
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < s.size; i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(strconv.Itoa(s.data[i]))
	}
	builder.WriteString("]")
	return builder.String()
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var set *Set
	var line, cmd string
	// v := NewSet(0)
	for scanner.Scan() {
		line = scanner.Text()
    	fmt.Println("$" + line)  
   		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			set = NewSet(cap)
		case "insert":
			for _, arg := range parts[1:] {
			value, _ := strconv.Atoi(arg)
			set.Insert(value)
			}
		case "show":
			fmt.Println(set.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !set.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(set.Contains(value))
		case "clear":
			set.size = 0
		}
	}
}
