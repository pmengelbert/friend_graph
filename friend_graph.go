package main

import "fmt"

type (
	Queue struct {
		size   int
		buffer []int
		front  int
		rear   int
	}
)

func NewQueue() *Queue {
	return &Queue{
		size:   4,
		buffer: make([]int, 5),
		front:  0,
		rear:   0,
	}
}

func (q *Queue) Insert(e int) {
	if q.rear == q.size {
		q.rear = 0
	} else {
		q.rear++
	}

	if q.rear == q.front {
		q.size *= 2
		newBuf := make([]int, q.size+1)

		latter := q.buffer[q.front:]
		copy(newBuf, latter)
		q.front = 0

		former := q.buffer[:q.rear]
		copy(newBuf[len(latter):], former)
		q.rear = len(latter) + len(former)

		q.buffer = newBuf
	}

	q.buffer[q.rear] = e
}

func (q *Queue) Remove() (int, error) {
	if q.front == q.rear {
		return -1, fmt.Errorf("underflow!")
	}

	if q.front == q.size {
		q.front = 0
	} else {
		q.front++
	}

	return q.buffer[q.front], nil
}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}

func main() {
	friends := [][]int{{1, 2}, {3, 4}, {3, 6}, {6, 7}, {1, 9}}
	fmt.Println("Input matrix:", friends)
	n := 7

	f1 := 4
	f2 := 7

	result, _ := isFriend(n, friends, f1, f2)
	fmt.Printf("Are %d and %d friends? -> %t\n", f1, f2, result)

	f1 = 3
	f2 = 6
	result, _ = isFriend(n, friends, f1, f2)
	fmt.Printf("Are %d and %d friends? -> %t\n", f1, f2, result)

	f1 = 1
	f2 = 6
	result, _ = isFriend(n, friends, f1, f2)
	fmt.Printf("Are %d and %d friends? -> %t\n", f1, f2, result)
}

func isFriend(n int, friends [][]int, f1, f2 int) (bool, error) {
	visitedOrAlreadyQueued := make(map[int]bool)

	graph := makeGraph(friends)

	// optionally print the graph by uncommenting the following:
	// printGraph(graph)

	queue := NewQueue()

	// execute the search using the queue for breadth-first

	var err error
	nodeIndex := f1
	queue.Insert(f1)

	for { // will exit when queue is empty
		if err != nil {
			return false, err
		}

		list := graph[nodeIndex]
		for _, val := range list {
			if val == f2 {
				return true, nil
			}

			if !visitedOrAlreadyQueued[val] {
				queue.Insert(val)
				visitedOrAlreadyQueued[val] = true
			}
		}

		if queue.isEmpty() {
			break
		}

		nodeIndex, err = queue.Remove()
	}

	return false, nil
}

// assumes there are no duplicate pairings in the input
// i.e. friends = [[1, 2], [1, 2]] shall not appear in the input
func makeGraph(friends [][]int) map[int][]int {
	graphBuilder := make(map[int][]int)
	for _, pair := range friends {
		f1 := pair[0]
		f2 := pair[1]

		graphBuilder[f1] = append(graphBuilder[f1], f2)
		graphBuilder[f2] = append(graphBuilder[f2], f1)
	}

	return graphBuilder
}
