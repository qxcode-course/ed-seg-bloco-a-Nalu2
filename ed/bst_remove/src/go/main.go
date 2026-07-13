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
	Left  *Node
	Right *Node
}

func insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{Value: val}
	}
	if val < node.Value{
		node.Left = insert(node.Left, val)
	} else if val > node.Value {
		node.Right = insert(node.Right, val)
	}
	return node
}

func BstInsert(values []int) *Node {
	var root *Node
	for _, val := range values {
		root = insert(root, val)
	}
	return root
}

func getMax(node *Node) *Node {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

func BstRemove(node *Node, value int) *Node {
	if node == nil {
		return nil
	} 
	if value < node.Value {
		node.Left = BstRemove(node.Left, value)
	} else if value > node.Value {
		node.Right = BstRemove(node.Right, value)
	}else {
		if node.Left == nil{
			return node.Right
		}
		if node.Right == nil{
			return node.Left
		}
		next := getMax(node.Left)
		node.Value = next.Value
		node.Left = BstRemove(node.Left, next.Value)
	}
	return node


}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	scanner.Scan()
	toRemove, _ := strconv.Atoi(scanner.Text())

	_ = toRemove // Ignora o valor a ser removido, pois não está implementado
	root := BstInsert(values)
	fmt.Println("original:")
	BShow(root, "") // Chama a função de impressão formatada
	root = BstRemove(root, toRemove)
	fmt.Println("modificado:")
	BShow(root, "") // Chama a função de impressão formatada da árvore modificada
}
