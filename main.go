package main

import (
	"algoritmos/queues"
	"algoritmos/stacks"
	"fmt"
)
import "algoritmos/quicksort"

func main() {
	fmt.Println("QuickSort")
	fmt.Println(quicksort.Run([]int{5, 6, 7, 2, 1, 0}))

	// Stack
	fmt.Println("Stack")
	stack := stacks.Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)

	// Queue
	fmt.Println("Queue")
	queue := queues.Queue{}
	queue.Enqueue(32)
	queue.Enqueue(64)
	queue.Enqueue(100)
	queue.Enqueue(3)
	fmt.Println(queue)
	queue.Denqueue()
	fmt.Println(queue)
	queue.Denqueue()
	fmt.Println(queue)
}
