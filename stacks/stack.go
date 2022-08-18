package stacks

type Stack struct{ itens []int }

func (stack *Stack) empty() bool {
	if stack.len() == 0 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) len() int {
	return len(stack.itens)
}

func (stack *Stack) Push(item int) {
	stack.itens = append(stack.itens, item)
}

func (stack *Stack) Pop() {
	if !stack.empty() {
		stack.itens = stack.itens[:stack.len()-1]
	}
}
