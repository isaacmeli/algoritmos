package stacks

type Stack struct {
	queue []int
	len   int
}

func (stack *Stack) empty() bool {
	if stack.len == 0 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) Push(item int) {
	stack.queue = append(stack.queue, item)
	stack.len = len(stack.queue)
}

func (stack *Stack) Pop() {
	if !stack.empty() {
		stack.len = stack.len - 1
		stack.queue = stack.queue[:stack.len]
	}
}
