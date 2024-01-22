package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)

}

func Pop(queue *list.List) interface{} {
	if queue.Len() == 0 {
		return nil
	}
	elem := queue.Front()
	queue.Remove(elem)
	return elem.Value
}

func printQueue(queue *list.List) {
	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		fmt.Print(elem.Value, "")
	}
}

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for e := l.Back(); e != nil; e = e.Prev() {
		reversedList.PushBack(e.Value)
	}
	return reversedList
}

const N = 7

func main() {

	mylist := list.New()

	for i := 0; i < N; i++ {
		Push(i, mylist)
	}

	for i := 0; i < 3; i++ {

		fmt.Println(Pop(mylist))

	}

	printQueue(mylist)

}
