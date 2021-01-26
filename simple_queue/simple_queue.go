package main

import "fmt"

type SimpleQueue struct {
	List []int
}

type EmptySliceError struct {
}

func (e *EmptySliceError) Error() string { return "Array is empty, nothing to dequeue." }

func (q *SimpleQueue) Enqueue(newItem int) {
	q.List = append(q.List, newItem)
}

func (q *SimpleQueue) Dequeue() (int, *EmptySliceError) {
	if len(q.List) == 0 {
		return 0, &EmptySliceError{}
	}
	lastItem := q.List[len(q.List)-1]
	q.List = q.List[:len(q.List)-1]
	return lastItem, nil
}

func main() {
	queue := SimpleQueue{make([]int, 0)}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	var res int
	var err *EmptySliceError
	for res, err = queue.Dequeue(); err == nil; res, err = queue.Dequeue() {
		fmt.Printf("%v\n", res)
	}
	fmt.Printf("%v\n", err)

}
