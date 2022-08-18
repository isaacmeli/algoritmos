package stacks

type Stack struct{ queue []int }

func (stack *Stack) empty() bool {
	if stack.len() == 0 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) len() int {
	return len(stack.queue)
}

func (stack *Stack) Push(item int) {
	stack.queue = append(stack.queue, item)
}

func (stack *Stack) Pop() {
	if !stack.empty() {
		stack.queue = stack.queue[:stack.len()-1]
	}
}
