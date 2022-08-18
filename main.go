package main

import (
	"algoritmos/stacks"
	"fmt"
)
import "algoritmos/quicksort"

func main() {
	fmt.Println(quicksort.Run([]int{5, 6, 7, 2, 1, 0}))

	stack := stacks.Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack)

	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
