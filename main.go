package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func Add[T any](head *List[T], value T) *List[T] {
	if head == nil {
		head = &List[T]{nil, value}
		return head
	}

	var last *List[T]
	for last = head; last.next != nil; last = last.next {
	}
	last.next = &List[T]{nil, value}

	return head
}

func PrintList[T any](list *List[T]) {
	for ; list != nil; list = list.next {
		fmt.Println(list.val)
	}
}

func main() {
	var list *List[int] = nil
	list = Add(list, 45)
	list = Add(list, 56)
	list = Add(list, 98)

	PrintList(list)
}
