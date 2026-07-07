package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root, size: 0}
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) PushFront(value int){
	newNode := &Node{value: value}
	firstNode := ll.root.next
	newNode.next = firstNode
	newNode.prev = ll.root

	ll.root.next = newNode
	firstNode.prev = newNode
	ll.size++
}

func (ll *LList) PushBack(value int){
	newNode := &Node{value: value}
	lastNode := ll.root.prev
	newNode.next = ll.root
	newNode.prev = lastNode

	lastNode.next = newNode
	ll.root.prev = newNode
	ll.size++
}

func (ll *LList) PopFront(){
	if ll.size == 0 {
		return
	}
	firstNode := ll.root.next
	nextToFirst := firstNode.next

	ll.root.next = nextToFirst
	nextToFirst.prev = ll.root

	ll.size--
}

func (ll *LList) PopBack() {
	if ll.size == 0 {
		return
	}
	lastNode := ll.root.prev
	prevToLast := lastNode.prev

	prevToLast.next = ll.root
	ll.root.prev = prevToLast

	ll.size--
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	curr := ll.root.next
	for curr != ll.root {
		sb.WriteString(strconv.Itoa(curr.value))
		if curr.next != ll.root {
			sb.WriteString(", ")
		}
		curr = curr.next
	}
	sb.WriteString("]")
	return sb.String()
}

func (ll *LList) Clear() {
    ll.root.next = ll.root
    ll.root.prev = ll.root
    ll.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for scanner.Scan() {

		line := scanner.Text()
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		fmt.Println("$" + line)
		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
