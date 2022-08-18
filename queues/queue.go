package queues

import "errors"

type Queue struct {
	itens []int
}

func (queue *Queue) len() int {
	return len(queue.itens)
}

func (queue *Queue) Enqueue(item int) {
	queue.itens = append(queue.itens, item)
}

func (queue *Queue) Denqueue() error {
	if queue.len() == 0 {
		return errors.New("not inicialized")
	}

	queue.itens = queue.itens[1:queue.len()]

	return nil
}
