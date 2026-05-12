package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	lastIndex := len(s.items) - 1
	lastValue := s.items[lastIndex]

	s.items = s.items[:lastIndex]

	return lastValue
}

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	firstValue := q.items[0]

	q.items = q.items[1:]

	return firstValue
}

func main2() {

	fmt.Println("STACK")

	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println()

	fmt.Println("QUEUE")

	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
