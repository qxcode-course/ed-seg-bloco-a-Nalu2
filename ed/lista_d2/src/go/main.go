package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node {}
	root.next = root
	root.prev = root
	root.root = root
	return &LList { root: root, size :0}
}

func (ll * LList) Size() int{
	return ll.size
}

func (ll * LList) Clear(){
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}
func (ll * LList) PushFront(value int) {
	newNode := &Node{Value: value, root: ll.root}
	firstNode := ll.root.next

	newNode.next = firstNode
	newNode.prev = ll.root

	ll.root.next = newNode
	firstNode.prev = newNode
	ll.size++
}

func (ll * LList) PushBack(value int) {
	newNode := &Node{Value: value, root: ll.root}
	lastNode := ll.root.prev

	newNode.next = ll.root
	newNode.prev = lastNode


	lastNode.next = newNode
	ll.root.prev = newNode
	ll.size++
}

func (ll *LList) PopFront() {
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

func (ll *LList) Front() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.prev
}

func (ll *LList) Search(value int) *Node {
	curr := ll.root.next
	for curr != ll.root {
		if curr.Value == value{
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}
	newNode := &Node{Value: value, root: ll.root}
	previousNode := node.prev

	newNode.next = node
	newNode.prev = previousNode

	previousNode.next = newNode
	node.prev = newNode
	ll.size++
}

func (ll *LList) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}
	nextNode := node.next
	previousNode := node.prev

	previousNode.next = nextNode
	nextNode.prev = previousNode
	ll.size--

	if nextNode == ll.root {
		return nil
	}
	return nextNode
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	curr := ll.root.next
	for curr != ll.root {
		sb.WriteString(strconv.Itoa(curr.Value))
		if curr.next != ll.root {
			sb.WriteString(", ")
		}
		curr = curr.next
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for scanner.Scan(){

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	node.Value = newvalue
			} else {
			 	fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Insert(node, newvalue)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Remove(node)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
