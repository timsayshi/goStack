package main

import "fmt"

type Node struct {
	Val  int
	Node *Node
}

type Stack []int

//isEmpty
func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

//push to stack
func (s *Stack) push(x int) {
	*s = append(*s, x)
}

//pop from stack
func (s *Stack) pop() (int, bool) {
	//if empty
	size := len(*s)
	if size == 0 {
		return 0, false
	}

	//pop from stack (top)
	element := (*s)[size-1]
	*s = (*s)[:size-1]
	return element, true
}

func main() {
	var stack Stack
	fmt.Println(stack.isEmpty())
	stack.push(4)
	stack.push(3)
	stack.push(2)
	stack.push(7)
	fmt.Println(stack.isEmpty())
	fmt.Println(stack)
	x, y := stack.pop()
	if y {
		fmt.Println("Popped", x, "from", stack)
	} else {
		fmt.Println("Stack is empty")
	}
}
